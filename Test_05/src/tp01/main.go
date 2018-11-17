package main

import (
	"fmt"
	"time"
)

//Go 通道

func main() {

	done := make(chan struct{})
	c := make(chan string)

	go func() {
		s := <-c
		fmt.Println("string:", s)
		close(done)
	}()
	time.Sleep(time.Second)
	c <- "hi!"
	<-done

}
