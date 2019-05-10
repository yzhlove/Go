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
	defer c.Close()
	values, err := redis.Values(c.Do("HSCAN", "User_Fragments:1", 0))
	if err != nil {
		panic(err)
	}
	result, _ := redis.Int(values[0], nil)
	fmt.Println(result)
	element, _ := redis.StringMap(values[1], nil)
	fmt.Println(element)

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
