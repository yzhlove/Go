package opt

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//ExampleStrTest string相关操作
func ExampleStrTest(client *redis.Client) {
	//0 表示不设置过期时间
	err := client.Set("string_name", "string_test", 0).Err()
	if err != nil {
		panic(err)
	}
	value, err := client.Get("string_name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string:name = ", value)
	//设置过期时间
	err = client.Expire("string_name", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	//设置过期时间
	err = client.Set("string_age", 18, 1*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	client.Incr("string_age")

	//获取
	value, err = client.Get("string_age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string_age_value = ", value)

	//incr
	client.Incr("string_age")

	value, err = client.Get("string_age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string_age_value = ", value)

	//decr
	client.Decr("string_age")

	value, err = client.Get("string_age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string_age_value = ", value)

	//显示过期时间
	ts, err := client.TTL("string_name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string_name TTL :" + ts.String())

	//显示过期时间
	ts, err = client.TTL("string_age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string_name TTL :" + ts.String())

	//getset
	lastValue, err := client.GetSet("string_name", "new_string_yzh").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string_name = " + lastValue)

	value, err = client.Get("string_name").Result()
	if err != nil {
		fmt.Println("string_name = " + value)
	}

	//关于空值
	fmt.Println("---------- nil -------------")
	client.Set("string_nil", "", 10*time.Second)

	nilvalue, err := client.Get("string_nil").Result()
	if nilvalue == "" {
		fmt.Println("string_nil == \"\" ")
	}

	//mset 与 mget
	err = client.MSet("key1", "love1", "key2", "love2").Err()
	if err != nil {
		panic(err)
	}
	valueArr, err := client.MGet("key1", "key2").Result()
	fmt.Printf("valueArr: %v \n", valueArr)


	
}
