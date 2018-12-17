package opt

import (
	"fmt"

	"github.com/go-redis/redis"
)

//ExampleHashTest hash操作
func ExampleHashTest(client *redis.Client) {

	people := make(map[string]interface{})
	people["name"] = "xjj"
	people["age"] = 18
	people["birthday"] = "1995-05-23"

	for k, v := range people {
		client.HSet("people", k, v)
	}
	//hmset
	// client.HMSet("people", people)

	//hgetall
	temp, err := client.HGetAll("people").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("people: %v \n", temp)

	//hlen
	length, _ := client.HLen("people").Result()
	fmt.Printf("people length = %v \n", length)

	//hget
	value, _ := client.HGet("people", "name").Result()
	fmt.Println("value name = ", value)

	//hmget
	keys := []string{"name", "age", "birthday"}
	tempArray, _ := client.HMGet("people", keys...).Result()
	fmt.Printf("tempArray: %v \n", tempArray)

	//hdel
	// client.HDel("people", []string{"age"}...).Err()

	//hkeys
	keyArray, _ := client.HKeys("people").Result()
	fmt.Printf("keys : %v \n", keyArray)

	//hvals
	valueArray, _ := client.HVals("people").Result()
	fmt.Printf("values: %v \n", valueArray)

	//hexists
	if client.HExists("people", "name").Val() {
		fmt.Println("name exists !")
	}

}
