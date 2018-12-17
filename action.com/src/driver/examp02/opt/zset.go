package opt

import (
	"fmt"

	"github.com/go-redis/redis"
)

//ExampleZSetTest zset示例
func ExampleZSetTest(client *redis.Client) {

	err := client.ZAdd("store_queue", redis.Z{Score: 1, Member: "one"}).Err()
	if err != nil {
		panic(err)
	}
	err = client.ZAdd("store_queue", redis.Z{Score: 2, Member: "two"}).Err()
	err = client.ZAdd("store_queue", redis.Z{Score: 3, Member: "three"}).Err()
	err = client.ZAdd("store_queue", redis.Z{Score: 4, Member: "four"}).Err()
	err = client.ZAdd("store_queue", redis.Z{Score: 5, Member: "five"}).Err()

	card := client.ZCard("store_queue").Val()
	fmt.Printf("card : %d \n", card)

	count, _ := client.ZCount("store_queue", "3", "5").Result()
	fmt.Printf("count : %d \n", count)

	//zrange
	result, _ := client.ZRange("store_queue", 0, -1).Result()
	fmt.Printf("result : %v \n", result)

	rs, err := client.ZRangeByScore("store_queue", redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	}).Result()
	fmt.Printf("rs : %v \n", rs)

	// => []z
	resultWithScore, err := client.ZRangeWithScores("store_queue", 0, -1).Result()
	fmt.Printf("resultWithScore: %v \n", resultWithScore)

	//zrank
	rank, err := client.ZRank("store_queue", "three").Result()
	fmt.Printf("three rank: %d \n", rank+1)

	//zscore
	score, err := client.ZScore("store_queue", "three").Result()
	fmt.Printf("three score : %.0f \n", score)

}
