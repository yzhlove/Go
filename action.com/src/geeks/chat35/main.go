package main

import (
	"fmt"
	"time"
)

//一个有趣的例子

func main() {

	stream := make(chan struct{})
	var counter int
	go func() {
		time.Sleep(5 * time.Second)
		close(stream)
	}()

loop:
	for {
		select {
		case <-stream:
			fmt.Println("stream running closed ...")
			break loop
		default:
			fmt.Println("default .")
		}
		counter++
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Done . counter = ", counter)
}
