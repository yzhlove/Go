package main

import (
	"driver/examp02/opt"
)

//redis 使用

func main() {

	client, err := opt.NewRedisClient()
	if err != nil {
		panic(err)
	}
	//string相关操作
	// opt.ExampleStrTest(client)

	//list相关操作
	// opt.ExampleListTest(client)

	//hash相关操作
	// opt.ExampleHashTest(client)

	//zset相关操作
	opt.ExampleZSetTest(client)

	client.Close()
}
