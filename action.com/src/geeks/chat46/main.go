package main

import "fmt"

//流处理

func main() {

	for _, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(mulit(add(v, 2), 2))
	}
	fmt.Println("Done .")
}

func mulit(value, tag int) int {
	return value * tag
}

func add(value, tag int) int {
	return value + tag
}
