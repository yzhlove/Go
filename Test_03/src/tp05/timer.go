package main

import (
	"fmt"
	"github.com/davyxu/cellnet/timer"
	"time"

	"github.com/davyxu/cellnet"
)

// 定时器测试
func main() {

	queue := cellnet.NewEventQueue()

	queue.StartLoop()

	var timeOut = 100

	timer.NewLoop(queue,time.Second , func(loop *timer.Loop) {

		timeOut--
		if 0 == timeOut {
			loop.Stop()
			queue.StopLoop()
		}

		fmt.Println("timeOut = ",timeOut)

	},nil).Start()

	queue.Wait()

}
