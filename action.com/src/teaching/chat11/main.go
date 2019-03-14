package main

import (
	"fmt"
	"time"
)

func main() {

	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i + 1
			time.Sleep(time.Second)
		}
	}()

	for data := range out {
		fmt.Println("data = ", data)
		if data == 10 {
			break
		}
	}

}
