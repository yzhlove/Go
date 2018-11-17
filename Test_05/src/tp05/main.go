package main

import (
	"fmt"
)

//通道检测

func main() {

	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-c
			if !ok { //OK => false 表示通道已经关闭
				return
			}
			fmt.Println("tmp ==> ", x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)

	<-done

}
