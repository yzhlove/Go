package main

import "fmt"

func main() {

	chanOwner := func() <-chan int {
		result := make(chan int, 5)
		go func() {
			defer close(result)
			for i := 0; i <= 5; i++ {
				result <- i
			}
		}()
		return result
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Println(result)
	}
	fmt.Println("Done .")
}
