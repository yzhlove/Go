package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	pool := NewRedisPool()
	c := pool.Get()

	if result, err := redis.String(c.Do("SET", "hello_world", "love xjj")); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

	_ = c.Close()
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
