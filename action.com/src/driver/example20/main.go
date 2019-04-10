package main

import (
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

//添加数据

func main() {
	pool := NewRedisPool()
	c := pool.Get()
	_, _ = c.Do("SET", "User_Info:1", "lcm")
	_, _ = c.Do("SET", "User_Info:2", "xjj")
	_, _ = c.Do("SET", "User_Info:3", "xyj")

	_, _ = c.Do("LPUSH", "User_List:1", "1001", "1002", "1003")
	_, _ = c.Do("LPUSH", "User_List:2", "2001", "2002", "2003")
	_, _ = c.Do("LPUSH", "User_List:3", "3001", "3002", "3003")

	_, _ = c.Do("HMSET", "User_Fragments:1", "name", "yzh", "age", "18")
	_, _ = c.Do("HMSET", "User_Fragments:2", "name", "xjj", "age", "19")
	_, _ = c.Do("HMSET", "User_Fragments:3", "name", "lcm", "age", "20")

	_, _ = c.Do("ZADD", "User_Items:1", "100", "yzh", "150", "xjj", "50", "lcm")
	_, _ = c.Do("ZADD", "User_Items:2", "200", "yzh", "20", "xjj", "2", "lcm")
	_, _ = c.Do("ZADD", "User_Items:3", "133", "yzh", "67", "xjj", "58", "lcm")

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
