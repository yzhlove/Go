package main

import (
	"fmt"
	"time"
)

//channel 函数参数

func worker(id int, c chan int) {
	for {
		fmt.Printf("id = %d worker value:%d \n", id, <-c)
	}
}

func main() {

	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
