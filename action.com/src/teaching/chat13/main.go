package main

import "fmt"

func main() {
	out := make(chan int, 3)
	fmt.Println(len(out))

	for i := 1; i <= 3; i++ {
		out <- i
	}

	fmt.Println(len(out))
}
