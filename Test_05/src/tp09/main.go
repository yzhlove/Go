package main

import (
	"fmt"
	"sync"
)

//通道 不允许逆序操作

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c //发送
	var recv <-chan int = c //接受

	go func() {
		defer wg.Done()

		for x := range recv {
			fmt.Println("x = ", x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()

}
