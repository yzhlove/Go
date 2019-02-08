package main

import (
	"fmt"
	"time"
)

//channel  的使用

func createWorker(id int) chan<- int {
	worker := make(chan int)
	go func() {
		for {
			fmt.Printf("id:%d worker value:%c \n", id, <-worker)
		}
	}()
	return worker
}

func channelDemo() {

	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 26; j++ {
			channels[i] <- 'A' + j
		}
	}
	time.Sleep(time.Second)
}

func main() {
	channelDemo()
}
