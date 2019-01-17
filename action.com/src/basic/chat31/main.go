package main

import (
	"fmt"
	"math/rand"
)

func main() {

	data := []string{"one", "two", "three", "four", "five"}

	for _, v := range data {
		fmt.Println("value:", v)
		//Fetch:
		for i := range []int{1, 2, 3, 4, 5} {
			fmt.Printf("%d ", i)
			if i == rand.Intn(5) {
				break
			}
		}
		fmt.Println(" Done")
	}

}
