package main

import (
	"fmt"
	"time"
)

//一个简单的协程池

func workerPool(n int, read chan int, write chan string) {
	for i := 0; i < n; i++ {
		go worker(i, read, write)
	}
}

func worker(id int, read chan int, write chan string) {
	for job := range read {
		result := fmt.Sprintf("worker %d process job: %d ", id, job)
		write <- result
	}
}

func genJob(n int) chan int {
	job := make(chan int, 200)
	go func() {
		for i := 0; i < n; i++ {
			job <- i
		}
		close(job)
	}()
	return job
}

func main() {

	job := genJob(200)
	read := make(chan string, 200)
	workerPool(5, job, read)
	time.Sleep(time.Second)
	close(read)
	for ret := range read {
		fmt.Println(ret)
	}
}
