package main

import (
	"fmt"
	"time"
)

func main() {

	bufferChan := make(chan int, 10)
	sign := make(chan byte, 2)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case bufferChan <- i:
			case bufferChan <- i + 10:
				//default:
				//	fmt.Println("Select Default")
			}
			//time.Sleep(time.Second)
		}
		close(bufferChan)
		fmt.Println("Close BufferChan .")
		sign <- 0
	}()

	go func() {
	loop:
		for {
			select {
			case data, ok := <-bufferChan:
				if !ok {
					fmt.Println("Close BufferChan .")
					break loop
				}
				fmt.Printf("data :%d \n", data)
				time.Sleep(2 * time.Second)
			}
		}
		sign <- 1
	}()
	<-sign
	<-sign
}
