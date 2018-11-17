package main

import (
	"fmt"
	"sync"
)

//通道不用，可以将已有通道设置为nil，这样select将不会选中此通道

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()
		for {
			select {
			case x, ok := <-a:
				if !ok {
					a = nil
					break
				}
				fmt.Println("a", x)
			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				fmt.Println("b", x)
			}
			if a == nil && b == nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()

}
