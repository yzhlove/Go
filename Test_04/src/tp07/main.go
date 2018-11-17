package main

import (
	"fmt"
	"time"
)

//异步执行

var index int

func counter() int {
	index++
	return index
}

func main() {

	num := 100

	go func(x, y int) {
		time.Sleep(1 * time.Second)
		fmt.Println(" go num = ", x, " index = ", y)
	}(num, counter())

	num += 100

	fmt.Println("num = ", num, " index = ", counter())

	time.Sleep(3 * time.Second)

}
