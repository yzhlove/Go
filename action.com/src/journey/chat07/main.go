package main

import (
	"fmt"
	"time"
)

//go协程初步

func channelDemo() {

	c := make(chan int)

	go func() {
		for {
			v := <-c
			fmt.Println("channel value:", v)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(1 * time.Millisecond)
}

func main() {
	channelDemo()
}
