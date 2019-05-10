package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	Max         int
	ConnTimeout time.Duration
	Host        string
	Index       int
	WaitTimeout time.Duration
}

func (c *RedisConfig) NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Max,
		IdleTimeout: c.ConnTimeout,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", c.Host, redis.DialDatabase(c.Index), redis.DialConnectTimeout(c.WaitTimeout))
		},
	}
}
