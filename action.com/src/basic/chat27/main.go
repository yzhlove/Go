package main

import (
	"fmt"
	"time"
)

//go channel 超时处理

func main() {

	cp := make(chan int, 1024)
	sign := make(chan int, 1)

	for i := 0; i < 1000; i++ {
		cp <- i
	}
	close(cp)

	go func() {
		var data int
		var ok = true

		for {
			select {
			case data, ok = <-cp:
				if !ok {
					fmt.Println("Read End ...")
					break
				}
				fmt.Printf("Read Data : %d \n", data)
			case ok = <-func() chan bool {
				timeout := make(chan bool, 1)
				go func() {
					fmt.Println("Write Before")
					time.Sleep(time.Millisecond)
					fmt.Println("Write After")
					timeout <- false
				}()
				return timeout
			}():
				fmt.Println("TimeOut ....")
				break
			}
			if !ok {
				close(sign)
				break
			}
		}
	}()
	<-sign
}
