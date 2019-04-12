package main

import (
	"log"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

//Migrate指令

const (
	SAVLE_HOST = "saved"
	SAVLE_PORT = "6379"
)

func main() {
	pool := NewRedisPool()
	c := pool.Get()
	defer c.Close()
	keys := []string{"User_Fragments:1", "User_Items:1", "User_Employees:1", "User_Info:1", "User_List:1"}
	for _, v := range keys {
		if result, err := redis.String(c.Do("MIGRATE", SAVLE_HOST, SAVLE_PORT, v, 0, 5000, "COPY")); err != nil {
			log.Printf("迁移错误:%v %v \n", v, err)
		} else {
			if result != "OK" {
				log.Printf("迁移失败:%v %v \n", v, err)
			} else {
				log.Printf("迁移成功:%v \n", v)
			}
		}
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
