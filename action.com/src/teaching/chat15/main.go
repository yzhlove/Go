package main

import (
	"fmt"
	"time"
)

func main() {

	exit := make(chan int)
	fmt.Println("start")

	index := 1

	time.AfterFunc(time.Second, func() {
		fmt.Printf("timer start : %d \n", index)
		exit <- 0
	})
	index = 2
	<-exit
}
