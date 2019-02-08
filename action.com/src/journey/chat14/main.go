package main

import (
	"fmt"
	"sync"
)

//waitgroup的使用

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorker(id int, wg *sync.WaitGroup) *worker {

	work := &worker{
		in: make(chan int),
		wg: wg,
	}
	go func() {
		for n := range work.in {
			fmt.Printf("id:%d value:%c \n", id, n)
			work.wg.Done()
		}
	}()
	return work
}

func channelDemo() {
	var workers [10]*worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = doWorker(i, &wg)
	}

	//wg.Add(20)
	for i, work := range workers {
		wg.Add(1)
		work.in <- i + 'A'

	}
	for i, work := range workers {
		wg.Add(1)
		work.in <- i + 'a'
	}
	wg.Wait()

}

func main() {
	channelDemo()
}
