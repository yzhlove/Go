package main

import (
	"fmt"
	"time"
)

func main() {

	intStream := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			<-intStream
			fmt.Println("value = ", i)
		}(i)
	}

	time.Sleep(time.Second)
	close(intStream)
	time.Sleep(time.Second)
	fmt.Println("Done .")
}
