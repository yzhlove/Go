package main

import (
	"fmt"
	"time"
)

func main() {

	out := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
	}()
	go func() {
		for {
			a, ok := <-out
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a = ", a)
		}
	}()
	fmt.Println("ok")
	time.Sleep(time.Second)
	close(out)
	time.Sleep(time.Second)
	fmt.Println("ok")
}
