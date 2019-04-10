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
	key := "User_Account:{1113008101029580800}"

	var result map[string]string

	result, _ = redis.StringMap(c.Do("HGETALL", key))
	for k, v := range result {
		fmt.Println(k, " -- ", v)
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
