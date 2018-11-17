package main

import "time"

// chan通道

func consumer(data chan int ,done chan bool) {

	for x := range data {
		println("recv:",x)
	}
	done <- true	//通知main ，消费结束
}


func producer(data chan int ) {
	for i := 0;i < 4;i++ {
		data <- i
		time.Sleep(time.Second)
	}
	close(data)
}


func main() {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data,done)
	go producer(data)

	<-done
}