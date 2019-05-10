package main

import (
	"fmt"
	"time"
)

// Go Select

func main() {

	cp := make(chan int, 1024)
	sign := make(chan int, 1)
	timeout := make(chan bool, 1)
	//向通道写入数据
	for i := 0; i < 1000; i++ {
		cp <- i
	}
	close(cp)

	go func() {
		time.Sleep(time.Millisecond)
		timeout <- false
	}()

	go func() {
		var d int
		var ok = true

		for {
			select {
			case d, ok = <-cp:
				if !ok {
					fmt.Println("Read End ...")
					break
				}
				fmt.Printf("Read Data : %d \n", d)
			case ok = <-timeout:
				fmt.Println("TimeOut ...")
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
