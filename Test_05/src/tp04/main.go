package main

import (
	"fmt"
)

//同步通道和异步通道的区别

//同步通道 len和cap 都返回 0

func main() {

	a, b := make(chan int), make(chan int, 3)

	b <- 1

	fmt.Println("a:", len(a), cap(a))
	fmt.Println("b:", len(b), cap(b))

}

/*
a: 0 0
b: 1 3
*/
