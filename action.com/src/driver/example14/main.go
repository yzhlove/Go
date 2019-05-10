package main

import (
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {

	pool := NewRedisPool()
	c := pool.Get()
	_, _ = c.Do("SADD", "User_Employees:1", "1001", "1002", "1003")
	_, _ = c.Do("SADD", "User_Employees:2", "2001", "2002", "2003")
	_, _ = c.Do("SADD", "User_Employees:3", "3001", "3002", "3003")
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
