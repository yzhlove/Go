package main

import (
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {

	pool := NewRedisPool()
	c := pool.Get()
	defer c.Close()

	key := "User_Employees:1"

	for i := 100; i < 1000; i++ {
		_, _ = c.Do("SADD", key, i)
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
