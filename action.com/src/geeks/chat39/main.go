package main

import (
	"fmt"
	"time"
)

//防止goroutine泄漏 为生产者提供一个退出标志

func main() {

	done := make(chan struct{})
	randStream := NewRandStream(done)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("time :", <-randStream)
	}
	close(done)
	time.Sleep(time.Second)
	fmt.Println("Done .")
}

func NewRandStream(done <-chan struct{}) <-chan int64 {
	randStream := make(chan int64)
	go func() {
		defer fmt.Println("new rand stream exit .")
		defer close(randStream)
		for {
			select {
			case randStream <- time.Now().UnixNano():
			case <-done:
				return
			}
		}
	}()
	return randStream
}
