package main

import (
	"fmt"
)

//切片的使用
func main() {

	ListArray := make([]string, 100)

	ListArray[1] = "Hello"
	ListArray[2] = "World"

	for k, v := range ListArray {
		fmt.Println("k = ", k, " v = ", v)
	}

}
