package main

import (
	"connection/driver"
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	rc := &driver.RedisConn{
		Addr:     "127.0.0.1:6379",
		PassWord: "",
		DB:       0,
	}

	dve := driver.NewClient(rc)
	if dve == nil {
		panic("NewClient Err!")
	}

	defer dve.(*redis.Client).Close()

	fmt.Printf("%T successful!", dve)

	result, err := dve.(*redis.Client).HGetAll("people").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("result = %v \n", result)

}
