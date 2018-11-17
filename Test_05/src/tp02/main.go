package main

import (
	"fmt"
)

//通道缓存区未写满或者数据数据未读取完毕，不会阻塞

func main() {

	c := make(chan int, 3) //创建带缓存的通道

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)

}
