package main

import (
	"fmt"
	"strconv"
)

func main() {

	hashMap := make(map[int][]string)

	for i := 0; i < 10; i++ {
		hashMap[0] = append(hashMap[0], strconv.Itoa(i))
	}

	fmt.Println(hashMap)

}
