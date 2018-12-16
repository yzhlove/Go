package main

import (
	"io"
	"log"
	"math/rand"
	pool "project/pool/util"
	"sync"
	"sync/atomic"
	"time"
)

//pool的使用

const (
	maxGorontines  = 100 //goruntinue数量
	pooledResource = 10  //池中的资源数量
)

//模拟需要存入进程池的资源
type dbConnection struct {
	ID int32
}

//Close 释放资源
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection ", dbConn.ID)
	return nil
}

//为每一个连接分配一个ID
var idCounter int32

//创建连接
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	var wg sync.WaitGroup
	wg.Add(maxGorontines)

	p, err := pool.New(createConnection, pooledResource)
	if err != nil {
		log.Println(err)
	}

	//使用池里的连接完成查询
	for query := 0; query < maxGorontines; query++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		go func(q int) {
			preformQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program.")
	p.Close()
}

//用来测试资源池
func preformQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	//将连接放回pool
	defer p.Release(conn)
	//模拟任务
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	log.Printf("----------- QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
