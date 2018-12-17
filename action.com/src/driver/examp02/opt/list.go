package opt

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//ExampleListTest list 相关操作
func ExampleListTest(client *redis.Client) {

	//队头添加
	client.LPush("fruit_queue", "apple")
	client.LPush("fruit_queue", "orange")
	client.LPush("fruit_queue", "banana")
	//队尾添加
	client.RPush("fruit_queue", "jujube")
	client.RPush("fruit_queue", "grape")

	length, err := client.LLen("fruit_queue").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("length = %v \n", length)

	//便利所有元素
	listArr, err := client.LRange("fruit_queue", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("list_array = %v \n", listArr)

	//返回index处的value
	value, _ := client.LIndex("fruit_queue", 2).Result()
	fmt.Println("index = 2 value = ", value)

	//给index位置处的元素赋值
	err = client.LSet("fruit_queue", 2, "love_fruit").Err()
	if err != nil {
		panic(err)
	}

	listArr, err = client.LRange("fruit_queue", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("list_array = %v \n", listArr)

	//rpoplpush (对同一个list操作 => 旋转操作)
	client.RPopLPush("fruit_queue", "fruit_queue").Err()

	listArr, err = client.LRange("fruit_queue", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("list_array = %v \n", listArr)

	//lpop队头弹出
	firstValue, err := client.LPop("fruit_queue").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("queueHead value = ", firstValue)

	//rpop队尾弹出
	lastValue, err := client.RPop("fruit_queue").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("queueTail value = ", lastValue)

	listArr, err = client.LRange("fruit_queue", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("list_array = %v \n", listArr)

	client.Expire("fruit_queue", 1*time.Second)

}
