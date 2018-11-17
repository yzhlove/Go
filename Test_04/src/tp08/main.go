package main

import (
	"fmt"
	"time"
)

//携程阻塞与退出

func main() {

	exit := make(chan struct{})

	go func() {
		defer close(exit)
		time.Sleep(3 * time.Second)
		fmt.Println("go run ...")
	}()

	fmt.Println("master thread run before ... ")

	<-exit //阻塞

	fmt.Println("master thred run after ...")

}
