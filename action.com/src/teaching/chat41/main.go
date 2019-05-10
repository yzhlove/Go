package main

import (
	"fmt"
	"runtime"
)

func consumer(ch chan string) {
	for {
		data := <-ch
		if data == "quit" {
			break
		}
		fmt.Printf("[Factory:%s ]\n ", data)
	}
	fmt.Println("[goroutine exit] ")
}

func main() {
	out := make(chan string)
	var (
		input string
		n     int
		err   error
	)
	for {

		if n, err = fmt.Scan(&input); err != nil {
			panic(err)
		}
		fmt.Printf("[input:%s count:%d ]\n", input, n)
		if n == 0 || input == "quit" {
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				out <- "quit"
			}
			continue
		}
		go consumer(out)
		fmt.Printf("[goroutines:%d ]\n", runtime.NumGoroutine())
	}
}
