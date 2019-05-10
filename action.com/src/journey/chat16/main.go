package main

import (
	"fmt"
	"math/rand"
	"time"
)

//select的使用

func generate() chan int {
	channel := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			channel <- i
			i++
		}
	}()
	return channel
}

func main() {
	var c1, c2 = generate(), generate()
	for {
		select {
		case n := <-c1:
			fmt.Println("c1 = ", n)
		case n := <-c2:
			fmt.Println("c2 = ", n)
		}
	}

}
