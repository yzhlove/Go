package main

import (
	"fmt"
	"time"
)

//缓冲通道与非缓冲通道的区别

func main() {

	bufferChan := make(chan int, 1)

	go func() {
		fmt.Println("Sleep Start ...")
		time.Sleep(time.Second)
		num := <-bufferChan
		fmt.Printf("Received %d \n", num)
	}()

	num := 1
	fmt.Printf("Send Message:%d \n", num)
	bufferChan <- num
	fmt.Println("Done.")
}
