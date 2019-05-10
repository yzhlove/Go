package main

import (
	"fmt"
	"time"
)

//channel 的使用

func worker(id int, aisle chan int) {
	for {
		fmt.Printf("worker id:%d value:%c \n", id, <-aisle)
	}
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Second)
}

func main() {
	chanDemo()
}
