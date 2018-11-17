package main

import (
	"fmt"
	"sync"
)

//即便是统一通道，select也是随机选择发送的

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer wg.Done()

		for {
			var v int
			var ok bool

			select {
			case v, ok = <-c:
				fmt.Println("a1:", v)
			case v, ok = <-c:
				fmt.Println("a2:", v)
			}

			if !ok { //如果通道关闭
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 10; i++ {
			select { //向同一通道发送select也是随机的
			case c <- i:
			case c <- i * 10:
			}
		}
	}()

	wg.Wait()
}
