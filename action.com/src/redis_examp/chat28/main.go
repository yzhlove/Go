package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {

	c := NewRedisPool().Get()
	defer c.Close()

	result, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
		panic(err)
	}

	quest := make(map[string]int)
	var str bytes.Buffer
	for _, v := range result {
		temp := strings.Split(v, ":")
		if _, ok := quest[temp[0]]; ok {
			continue
		}
		str.WriteString("\"" + v + "\",\n")
		quest[temp[0]] = 1
	}
	fmt.Println(str.String())
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
