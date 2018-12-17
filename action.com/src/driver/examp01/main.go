package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//go操作redis

func main() {

	client := NewClient()
	OptRedis(client)

}

//NewClient 新建客户端
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping().Result()
	fmt.Println(ping, err)
	return client
}

//OptRedis 操作redis
func OptRedis(client *redis.Client) {

	defer client.Close()

	key := "string_test"

	err := client.Set(key, "Hello World", 1*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	value, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key:" + key + " value:" + value)

}
