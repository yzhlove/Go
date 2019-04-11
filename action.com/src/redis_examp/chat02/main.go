package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

//分段读取list

const MAX = 200

func main() {
	pool := NewRedisPool()
	c := pool.Get()
	defer c.Close()
	key := "User_List:3"
	//分批读取
	length, _ := redis.Int(c.Do("LLEN", key))
	fmt.Println("length = ", length)

	values, _ := redis.Strings(c.Do("LRANGE", key, 0, 200))
	fmt.Println(values)
	fmt.Println(len(values))

	index := length / 30
	if length%30 != 0 {
		index += 1
	}

	for i := 0; i <= index; i++ {
		first := i * 30
		last := (i+1)*30 - 1
		values, _ := redis.Strings(c.Do("LRANGE", key, first, last))
		fmt.Println("===========================================")
		fmt.Println(len(values), " ", first, " ", last)
		fmt.Println(values)

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
