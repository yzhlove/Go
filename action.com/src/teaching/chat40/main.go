package main

import (
	"fmt"
	"runtime"
)

//协程不退出，危险

func consumer(ch chan string) {
	for {
		d := <-ch
		fmt.Println("data = ", d)
	}
}

func main() {
	out := make(chan string)

	for {
		var str string
		_, _ = fmt.Scan(&str)
		go consumer(out)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}
}
