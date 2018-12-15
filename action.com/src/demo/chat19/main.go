package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//无缓冲通道的使用

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	//创建一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)

	go player("Tom", court)
	go player("Jerry", court)

	//发球 写入数据
	court <- 1

	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s Won \n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n", name)
			//关闭通道
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++
		court <- ball
	}
}
