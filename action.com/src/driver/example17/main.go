package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

//list设置

func main() {
	pool := NewRedisPool()
	c := pool.Get()
	key := "User_List:1"
	//_, _ = redis.Int(c.Do("LPUSH", key, "a", "b", "c", "d", "e"))

	values, _ := redis.Values(c.Do("LRANGE", key, 0, -1))

	key2 := "User_List:2"

	for i := 0; i < len(values); i++ {
		fmt.Println(string(values[i].([]byte)))
		_, _ = c.Do("RPUSH", key2, string(values[i].([]byte)))
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
