package db

import (
	"fmt"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

type ArgsOption struct {
	MaxThreads int
	Host       string
}

func (opt *ArgsOption) NewConnPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     opt.MaxThreads,
		IdleTimeout: time.Duration(5 * time.Second),
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", opt.Host, redis.DialDatabase(0), redis.DialConnectTimeout(time.Duration(2*time.Second)))
		},
	}
}

//master pool
func GetRequestPool() (pool *redis.Pool, err error) {
	opt := ArgsOption{MaxThreads: runtime.NumCPU(), Host: "127.0.0.1:6379"}
	if pool = opt.NewConnPool(); pool == nil {
		return nil, fmt.Errorf("[ERROR] get redis pool error:%v \n", err)
	}
	if err = Ping(pool); err != nil {
		return nil, fmt.Errorf("[ERROR] ping redis err:%v \n", err)
	}
	return
}

//saved pool
func GetResponsePool() (pool *redis.Pool, err error) {
	opt := ArgsOption{MaxThreads: runtime.NumCPU(), Host: "127.0.0.1:6380"}
	if pool = opt.NewConnPool(); pool == nil {
		return nil, fmt.Errorf("[ERROR] get redis pool error:%v \n", err)
	}
	if err = Ping(pool); err != nil {
		return nil, fmt.Errorf("[ERROR] ping redis err:%v \n", err)
	}
	return
}

func Ping(pool *redis.Pool) error {
	var (
		conn = pool.Get()
		err  error
	)
	if err = conn.Err(); err != nil {
		return err
	}
	if _, err = conn.Do("PING"); err != nil {
		return err
	}
	_ = conn.Close()
	return nil
}
