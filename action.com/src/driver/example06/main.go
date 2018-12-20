package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
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

		client := NewClient()
		key := "hashMap"
		initData(key,client)

		//dataMap := getData(key,1,10,client)
		//fmt.Printf("dataMap = %v \n",dataMap)

		getPageData(key,0,-1,client)

		//[]redis.Z [{60 hxy} {80 lcm} {100 xyj} {110 yzh} {120 xjj} {145 lyf} {150 fyb} {180 lz} {200 gakki} {2000 yurisa}]

		client.Expire(key,1 * time.Second)
		client.Close()
}

func initData(dataKey string,client *redis.Client) {
	dataMap := map[string]int {
		"yzh":110,
		"xjj":120,
		"xyj":100,
		"lcm":80,
		"fyb":150,
		"hxy":60,
		"lyf":145,
		"lz":180,
		"gakki":200,
		"yurisa":2000,
	}

	for key , value := range dataMap {
		client.ZAdd(dataKey,redis.Z{
			Member:key,
			Score:float64(value),
		})
	}
}


func getData(key string , page int ,count int , client *redis.Client) map[string]float64 {

	number := client.ZCard(key).Val()
	fmt.Printf("number = %v \n",number)
	dataMap := make(map[string]float64)
	redisData , err := client.ZRangeWithScores(key,0 ,-1).Result()
	if err != nil {
		panic(err)
	}
	for _ , value := range redisData {
		dataMap[value.Member.(string)] = value.Score
	}
	return dataMap
}

func getPageData(key string, page int , count int64 ,client *redis.Client) {

	start := int64(page - 1) * count
	end := int64(page) * count - 1

	redisData , err := client.ZRangeWithScores(key,start,end).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T %v \n",redisData,redisData)
}