package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Print("i:", i, " ")
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Print("->:", i, " ")
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Second)
}
