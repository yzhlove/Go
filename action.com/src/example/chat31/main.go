package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.Once的使用

var once sync.Once

func main() {

	words := []string{"I", "Love", "You"}

	for i, v := range words {
		once.Do(onces)
		fmt.Println("count = ", i, " world = ", v)
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(8000)
}

func onces() {
	fmt.Println(" onces ")
}

func onced() {
	fmt.Println(" onced ")
}
