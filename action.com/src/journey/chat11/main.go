package main

import (
	"fmt"
	"time"
)

//buffer channel

func worker(id int, work chan int) {

	//第一种 判断channel是否关闭
	for nt := range work {
		//第二种 判断channel是否关闭
		if n, ok := <-work; !ok {
			break
		} else {
			fmt.Printf("id:%d worker:%d \n", id, n)
		}
		fmt.Printf("id:%d worker:%d \n", id, nt)
	}
}

func bufferChannelDemo() {
	job := make(chan int, 3)
	job <- 1
	job <- 2
	job <- 3
	close(job)
	go worker(0, job)

	time.Sleep(time.Second)
}

func main() {
	bufferChannelDemo()
}
