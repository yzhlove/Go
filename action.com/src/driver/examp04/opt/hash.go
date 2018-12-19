package opt

import (
	"driver/examp04/ssdb"
	"fmt"
)

//ExampleHash hasn例子
func ExampleHash(db *ssdb.Client) {

	result, _ := db.Do("hset", "hash_map_value", "name", "yzh", "birthday", "1996-12-24")
	fmt.Printf("result = %v \n", result)

	result, _ = db.Do("hset", "hash_map_value", "birthday", "1996-12-24")

	result, _ = db.Do("hgetall", "hash_map_value")
	fmt.Printf("result = %v \n", result)

	result, _ = db.Do("hget", "hash_map_value", "birthday")
	fmt.Printf("result = %v \n", result)

	db.Do("hset", "hash_map_value", "sex", "man")

	result, _ = db.Do("hget", "hash_map_value", "sex")
	fmt.Printf("result = %v \n", result)

	result, _ = db.Do("hdel", "hash_map_value", "sex")
	fmt.Printf("del result = %v \n", result)

	result, _ = db.Do("hget", "hash_map_value", "sex")
	fmt.Printf("get del type result = %v \n", result)

	//multi_hset
	result, _ = db.Do("multi_hset", "people", "name", "xjj", "birthday", "1996-05-23", "sex", "woman")
	fmt.Printf("result: %v \n", result)

	//multi_hget
	result, _ = db.Do("multi_hget", "people", "name", "sex")
	fmt.Printf("result = %v \n", result)

	result, _ = db.Do("multi_hget", "people", "birthday", "name")
	fmt.Printf("result = %v \n", result)

	

}
