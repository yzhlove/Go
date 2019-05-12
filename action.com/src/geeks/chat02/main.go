package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
)

func main() {

	fmt.Println("BEFORE:", runtime.NumGoroutine())
	fmt.Println(FirstResp())
	fmt.Println("AFTER:", runtime.NumGoroutine())
}

func GetResp(i int) string {
	temp := []string{"xjj", "xyj", "fyb", "lcm"}
	return strconv.Itoa(i) + ":" + temp[rand.Intn(len(temp))]
}

func FirstResp() string {
	index := 12
	out := make(chan string, index)
	//out := make(chan string)
	for i := 0; i < index; i++ {
		go func(i int) {
			out <- GetResp(i)
		}(i)
	}
	return <-out
}
