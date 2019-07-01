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

	key := "User_Equipments:{1135436627158503425}"
	//_, _ = c.Do("ZADD", key, "5", "e", "6", "f", "7", "g", "8", "h")
	var result []string
	var err error
	if result, err = redis.Strings(c.Do("SMEMBERS", key)); err != nil {
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
