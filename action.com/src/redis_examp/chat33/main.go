package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

//redis connection timeout err

var (
	wg = new(sync.WaitGroup)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pool := GetPool()
	c := pool.Get()
	defer c.Close()
	count := 1000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go Run(i, c, wg)
	}

	wg.Wait()
	fmt.Println("Done .")
}

func GetPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: time.Duration(2 * time.Second),
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(0), redis.DialConnectTimeout(time.Duration(2*time.Second)))
		},
	}
}

func Run(id int, c redis.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	//c := pool.Get()
	//defer c.Close()
	if err := ping(c); err != nil {
		log.Printf("[ERROR] ping err : %v \n", err)
		return
	}
	ts := rand.Intn(5) + 3
	fmt.Println("[INFO] before time ", ts)
	time.Sleep(time.Duration(ts) * time.Second)
	fmt.Printf("[%v] successful .\n", id)
}

func ping(c redis.Conn) (err error) {
	if err = c.Err(); err != nil {
		return
	}
	_, err = c.Do("PING")
	return
}
