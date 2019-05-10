package main

import (
	"fmt"
	"time"
)

//go协程使用

func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello Worker: ", i)
			}
		}(i)
	}
	time.Sleep(1 * time.Millisecond)
}
