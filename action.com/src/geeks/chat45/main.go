package main

import "fmt"

//一个无聊的pipeline

//批处理

func main() {

	for _, v := range add(mulitply([]int{1, 2, 3}, 2), 5) {
		fmt.Println("value = ", v)
	}
	fmt.Println("Done .")
}

func mulitply(values []int, tag int) []int {
	temp := make([]int, len(values))
	for i, v := range values {
		temp[i] = tag * v
	}
	return temp
}

func add(values []int, tag int) []int {
	temp := make([]int, len(values))
	for i, v := range values {
		temp[i] = v + tag
	}
	return temp
}
