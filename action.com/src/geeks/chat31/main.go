package main

import (
	"fmt"
	"strconv"
	"time"
)

//带缓冲的channel

func main() {

	stringStream := make(chan string, 10)

	go func() {
		for i := 0; i < 20; i++ {
			value, ok := <-stringStream
			fmt.Printf("status:%v value:%v \n", ok, value)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		stringStream <- strconv.Itoa(i + 1)
	}
	fmt.Println("waiting ... ")
	time.Sleep(7 * time.Second)
	fmt.Println("close ... ")
	close(stringStream)

	time.Sleep(time.Minute)

}
