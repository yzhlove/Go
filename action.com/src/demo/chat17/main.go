package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//Store / Load 的使用

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {

	wg.Add(2)

	go doWork("A")
	go doWork("B")

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(500 * time.Millisecond)
		}
		//设置值，当值为1的时候，协程工作停止
		atomic.StoreInt64(&shutdown, 1)
	}()

	wg.Wait()

}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s is work\n", name)
		time.Sleep(250 * time.Millisecond)
		//检测任务停止的标志
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
