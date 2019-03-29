package main

import "fmt"

func main() {

	var list []int

	for i := 0; i < 10; i++ {
		list = append(list, i)
	}

	fmt.Printf("%v \n", list)

	list2 := make([]int, 10)

	for i := 0; i < len(list2); i++ {
		list2[i] = i
	}

	fmt.Printf("%v \n", list2)
}
