package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//带缓冲通道的使用

const (
	numberGoroutines = 4  //goroutine 数量
	taskLoad         = 10 //task 数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	tasks := make(chan string, taskLoad)

	//启动任务组
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//投递任务
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task:%d ", post)
	}
	//当所有任务都完成的时候关闭通道,以便goroutine退出
	close(tasks)
	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("通道关闭,没有任务了! woker[%d]退出\n", worker)
			return
		}

		sleep := rand.Int63n(1000) + 500
		fmt.Printf("Woker:%d Task:%s 开始工作! time:[%d]\n", worker, task, sleep)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker:%d TaskId:%s 完成任务!\n", worker, task)
	}
}
