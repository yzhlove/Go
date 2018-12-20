package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password:"",
		DB:0,
	})
	_ , err := client.Ping().Result()
	if err != nil {
		panic(err)
		return nil
	}
	return client
}

func main() {

	redisClient := NewClient()
	if redisClient == nil {
		fmt.Println("redis init Err")
		return
	}

	key := "string_name"
	result := redisClient.Exists(key).Val()
	fmt.Println("result = " , result )

	key = "finish_rank_data"
	flag := redisClient.ZIncrBy(key,float64(10),"hello").Val()
	fmt.Println(flag )


}

