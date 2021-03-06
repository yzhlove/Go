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
	pool.Wait = true
	count := 500
	wg.Add(count)
	for i := 0; i < count; i++ {
		go Run(i, pool, wg)
	}

	wg.Wait()
	fmt.Println("Done .")
}

func GetPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     4,
		IdleTimeout: time.Duration(10 * time.Second),
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(0), redis.DialConnectTimeout(time.Duration(10*time.Second)))
		},
	}
}

func Run(id int, pool *redis.Pool, wg *sync.WaitGroup) {
	defer wg.Done()
	c := pool.Get()
	defer c.Close()
	if err := ping(c); err != nil {
		log.Printf("[ERROR] ping err : %v \n", err)
		return
	}
	ts := rand.Intn(5) + 3
	fmt.Println("[INFO] before time ", ts)
	time.Sleep(time.Duration(ts) * time.Second)
	if err := ping(c); err != nil {
		log.Printf("[ERROR] ping err : %v \n", err)
		return
	}
	fmt.Printf("[%v] successful .\n", id)
}

func ping(c redis.Conn) (err error) {
	if err = c.Err(); err != nil {
		return fmt.Errorf(" Connection Error .. %v ", err)
	}
	_, err = c.Do("PING")
	return
}
