package main

import "log"

// 多次补货，只有一次成功

func catch() {
	log.Println("catch:",recover())
}

func main() {
	defer catch()					// 成功
	defer log.Println(recover())	// 失败
	defer println(recover())					// 失败

	panic("what are you going")
}
