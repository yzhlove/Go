package main

import "fmt"

func main() {
	str := make([]int, 0, 5)
	str = append(str, 1, 2, 3, 4, 5)
	fmt.Printf("%v ", str)
}
