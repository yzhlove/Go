package main

import (
	"runtime"
	"time"

	"fmt"

	"github.com/garyburd/redigo/redis"
)

//redis scan使用

func main() {

	pool := NewRedisPool()
	c := pool.Get()
	values, _ := redis.Values(c.Do("SCAN", 0))

	fmt.Printf("%T %v \n", values, values)

	fmt.Println(string(values[0].([]byte)))

	for _, value := range values[1].([]interface{}) {
		fmt.Println(string(value.([]byte)))
	}

}

func NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2 * runtime.NumCPU(),
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379", redis.DialDatabase(0), redis.DialConnectTimeout(2*time.Second))
		},
	}
}
