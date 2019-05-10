package main

import (
	"fmt"
	"strconv"
)

//指定容量的map

func main() {

	data := make(map[int]string, 3)
	for i := 1; i <= 20; i++ {
		tmp := strconv.Itoa(i)
		data[i] = tmp
		fmt.Printf("data-key:%d data-value:%v len:%d \n", i, data[i], len(data))
	}

}
