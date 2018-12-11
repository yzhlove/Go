package main

import (
	"fmt"
)

func main() {

	list := make([]int, 0, 4)

	for i := 1; i <= cap(list); i++ {
		list = append(list, i)
	}
	fmt.Printf("list: %v \n", list)

	length := len(list)
	for i := 0; i < length; i++ {

		fmt.Printf("length = %v \n", len(list))
		list = append(list[:0], list[1:]...)
	}

}
