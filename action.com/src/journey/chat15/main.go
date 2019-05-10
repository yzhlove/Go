package main

import (
	"fmt"
	"sync"
)

//channel 的使用

type worker struct {
	in       chan int
	callback func()
}

func createWorker(id int, callback func()) *worker {
	work := new(worker)
	work.in = make(chan int)
	work.callback = callback
	go func() {
		for n := range work.in {
			fmt.Printf("id %d value:%c \n", id, n)
			work.callback()
		}
	}()
	return work
}

func channelDemo() {
	var workers [10]*worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, func() {
			wg.Done()
		})
	}

	for i, work := range workers {
		work.in <- 'A' + i
	}

	for i, work := range workers {
		work.in <- 'a' + i
	}

	wg.Wait()

}

func main() {
	channelDemo()
}
