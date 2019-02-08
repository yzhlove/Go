package main

import "fmt"

//select的使用

type worker struct {
	in     chan int
	status chan bool
}

func doWorker(id int) worker {

	work := worker{
		in:     make(chan int, 3),
		status: make(chan bool),
	}
	go func() {
		for n := range work.in {
			fmt.Printf("id:%d , value:%c \n", id, n)
			work.status <- true
		}
	}()
	return work
}

func channelDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = doWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].status
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'K' + i
		<-workers[i].status
	}

}

func main() {
	channelDemo()
}
