package main

import (
	"fmt"
)

/*
通道规则：
* 向已经关闭的通道发送数据，引发panic
* 从已关闭的通道接受数据，返回缓存的值或这0
* 无论收发，nil通道都会阻塞
* 重复关闭，或者关闭nil通道会引发panic
*/

func main() {

	c := make(chan int, 3)
	c <- 10
	c <- 20
	close(c)
	//通道关闭之后接受的都是缓存数据
	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		fmt.Printf("%v , %v ,%v \n", i, ok, x)
	}
}
