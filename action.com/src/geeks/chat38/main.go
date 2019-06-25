package main

import (
	"fmt"
	"time"
)

//goroutine 退出

func main() {

	done := make(chan struct{})
	tm := work(done, nil)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("cancel work...")
		close(done)
	}()

	<-tm
	fmt.Println("Done .")

}

func work(done <-chan struct{}, data <-chan int) <-chan int {
	tm := make(chan int)
	go func() {
		defer fmt.Println("exit.")
		defer close(tm)
		for {
			select {
			case s := <-data:
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return tm
}
