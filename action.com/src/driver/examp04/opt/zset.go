package opt

import (
	"driver/examp04/ssdb"
	"fmt"
)

//ExampleZset zset操作
func ExampleZset(db *ssdb.Client) {

	db.Do("zset", "link_queue", "goolge", 100)
	db.Do("zset", "link_queue", "microsoft", 90)
	db.Do("zset", "link_queue", "baidu", 70)
	db.Do("zset", "link_queue", "ali", 77)
	db.Do("zset", "link_queue", "telnet", 76)

	result, _ := db.Do("zrange", "link_queue", 0, -1)
	fmt.Printf("result = %v \n", result)

}
