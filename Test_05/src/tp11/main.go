package main

import (
	"fmt"
	"sync"
)

//通道收发

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name, x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 20; i++ {
			select { //select 随机选择一个通道接受数据
			case a <- i:
			case b <- i:
			}
		}
	}()

	wg.Wait()

}
