package main

import (
	"fmt"
	"time"
)

//通道的便利

func main() {

	passage := make(chan int)
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	go func() {
		for _, value := range arr {
			passage <- value
			time.Sleep(200 * time.Millisecond)
		}
		close(passage)
	}()

	for v := range passage {
		fmt.Printf("value = %v \n", v)
	}
}
