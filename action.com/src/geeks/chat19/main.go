package main

import (
	"sync"
	"time"

	"fmt"
)

//go cond

func main() {

	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0, 10)
	var removeQueue = func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		fmt.Println("remove queue = ", queue[0])
		queue = queue[1:]
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 1; i <= 10; i++ {
		cond.L.Lock()
		for len(queue) == 2 {
			cond.Wait()
		}
		queue = append(queue, i)
		fmt.Println("add value ", i)
		go removeQueue(1 * time.Second)
		cond.L.Unlock()
	}

}
