package main

import (
	"fmt"
)

//通道使用 案例一

func main() {

	cp := make(chan int, 1000)
	notify := make(chan int, 1)

	//write chan
	for i := 0; i < 1000; i++ {
		cp <- i
	}
	close(cp)

	go func() {
		var e int
		var ok = true
		for {
			select {
			case e, ok = <-cp:
				if !ok {
					fmt.Println("Read Exit")
					//close(notify)
					notify <- 0
					return
				}
				fmt.Printf("Read Data :%d \n", e)
			}
		}
	}()

	<-notify
}
