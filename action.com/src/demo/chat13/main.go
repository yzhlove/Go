package main

import (
	"fmt"
	"runtime"
	"sync"
)

//goroutine 调度

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Exiting ...")

}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 20000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d \n", prefix, outer)
	}
	fmt.Println("Completed ", prefix)
}
