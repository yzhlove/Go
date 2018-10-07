package main

import "fmt"

// 可变数组 (切片)

func main() {

	x := [][]int{
		{1,2},
		{3,4,5},
		{6},
	}

	fmt.Println(x[1])

	x[2] = append(x[2],7,8,9,10)

	fmt.Println(x[2])

}
