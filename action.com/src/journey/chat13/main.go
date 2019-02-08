package main

import "fmt"

//channel 的使用

type worker struct {
	in   chan int
	flag chan bool
}

func doWorker(id int) worker {

	work := worker{
		in:   make(chan int),
		flag: make(chan bool),
	}
	go func() {
		for n := range work.in {
			fmt.Printf("id:%d value:%c \n", id, n)
			go func() { work.flag <- true }()
		}
	}()
	return work
}

func channelDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = doWorker(i)
	}

	for i, work := range workers {
		work.in <- i + 'A'
	}

	for i, work := range workers {
		work.in <- i + 'a'
	}

	for _, work := range workers {
		<-work.flag
		<-work.flag
	}
}

func main() {
	channelDemo()
}
