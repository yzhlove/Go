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
	key := "zset_key:1"
	//_, _ = c.Do("ZADD", key, "5", "e", "6", "f", "7", "g", "8", "h")
	var result map[string]int64
	var err error
	if result, err = redis.Int64Map(c.Do("ZRANGE", key, 0, -1, "WITHSCORES")); err != nil {
		panic(err)
	}
	for index, value := range result {
		fmt.Println(index, " <> ", value)
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
