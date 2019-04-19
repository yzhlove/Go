package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	in := request()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(in chan int, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			count := 0
			for v := range in {
				fmt.Printf("[%d ] value = %v \n", i, v)
				count++
			}
			fmt.Printf("[%d ] count =  %v \n", i, count)
		}(in, i, &wg)
	}
	wg.Wait()
	fmt.Printf("ok .")
}

func request() chan int {
	in := make(chan int, 10)
	go func() {
		for i := 1; i < 100; i++ {
			in <- i
		}
		time.Sleep(1 * time.Second)
		close(in)
	}()
	return in
}
