package main

import "fmt"

func main() {

	out := make(chan int, 2)

	for i := 1; i <= 2; i++ {
		out <- i
	}

	close(out)

	for i := 0; i < cap(out)+5; i++ {
		v, ok := <-out
		fmt.Println(v, ok)
	}
}
