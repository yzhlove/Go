package db

import (
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

func NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2 * runtime.NumCPU(),
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379", redis.DialDatabase(0), redis.DialConnectTimeout(2*time.Second))
		},
	}
}

func GetKeys() []string {

	c := NewRedisPool().Get()
	defer c.Close()

	result, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
		panic(err)
	}

	return result
}
