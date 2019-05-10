package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	out := make(chan int, 100)
	for i := 0; i < 10; i++ {
		out <- i
	}

	go func() {
		index := 0
		for {
			a, ok := <-out
			if !ok {
				os.Exit(0)
			}
			fmt.Println("out = ", a, " index = ", index)
			index++
		}
	}()

	close(out)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}
