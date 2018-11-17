package main

import (
	"fmt"
	"runtime"
	"time"
)

//go 终止调用栈

func main() {

	for i := 0; i < 2; i++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c : %d \n", 'a'+x, n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}
	runtime.Goexit()
	fmt.Println("main.exit ...")
}
