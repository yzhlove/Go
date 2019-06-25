package main

import (
	"fmt"
	"time"
)

//限制

func main() {

	result := func() <-chan int {
		out := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				out <- i + 1
			}
			time.AfterFunc(time.Second, func() {
				close(out)
			})
		}()
		return out
	}

	read := func(stream <-chan int) {
		for value := range stream {
			fmt.Println("value = ", value)
		}
	}

	temp := result()
	read(temp)

	fmt.Println("Done .")
}
