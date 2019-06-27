package main

import (
	"fmt"
)

//pipeline

func main() {

	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case intStream <- i:
				case <-done:
					return
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, tag int) <-chan int {
		multiStream := make(chan int)
		go func() {
			defer close(multiStream)
			for i := range intStream {
				select {
				case multiStream <- i * tag:
				case <-done:
					return
				}
			}
		}()
		return multiStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, tag int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range intStream {
				select {
				case addStream <- i + tag:
				case <-done:
					return
				}
			}
		}()
		return addStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println("value = ", v)
	}

	fmt.Println("Done .")

}
