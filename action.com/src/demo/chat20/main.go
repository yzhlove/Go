package main

import (
	"fmt"
	"sync"
	"time"
)

//接力比赛

var wg sync.WaitGroup

func main() {

	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()
}

//Runner 奔跑
func Runner(baton chan int) {

	var newRunner int

	//等待接力棒
	runner := <-baton

	fmt.Printf("Runner %d Running With Both\n", runner)

	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished ,Race Over \n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d \n", runner, newRunner)
	//交接
	baton <- newRunner
}
