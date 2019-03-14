package main

import "fmt"

func printer(input chan int) {
	for {
		data := <-input
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	input <- 0
}

func main() {

	out := make(chan int)
	go printer(out)
	for i := 1; i <= 10; i++ {
		out <- i
	}
	out <- 0
	<-out
}
