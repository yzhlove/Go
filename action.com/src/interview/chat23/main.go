package main

import (
	"fmt"
	"time"
)

//任务取消

func main() {

	temp := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, temp chan struct{}) {
			for {
				if IsCancel(temp) {
					break
				}
				time.Sleep(time.Second * 5)
			}
			fmt.Println(i, "->cancel")
		}(i, temp)
	}

	//cancel2(temp)
	//time.Sleep(1 * time.Second)

	cancel1(temp)
	time.Sleep(1 * time.Second)
}

func IsCancel(c chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

func cancel1(c chan struct{}) {
	c <- struct{}{}
}

func cancel2(c chan struct{}) {
	close(c)
}
