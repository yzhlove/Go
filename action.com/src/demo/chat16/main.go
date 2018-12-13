package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//竞态资源加锁

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {

	wg.Add(2)

	go inCounter(1)
	go inCounter(2)

	wg.Wait()

	fmt.Printf("counter: %v \n", counter)

}

func inCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
