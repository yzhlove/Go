package main

import "fmt"

func main() {

	out := make(chan int)
	close(out)
	fmt.Printf("ptr:%p cap:%d len:%d \n", out, cap(out), len(out))

	out <- 1

}
