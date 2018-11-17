package main

import (
	"fmt"
	"time"
)

//　当所有通道不可用的时候，select会选择执行default

func main() {

	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data_c:", x)
			default:
			}
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	c <- 100
	close(c)

	<-done
}
