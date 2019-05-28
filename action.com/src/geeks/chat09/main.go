package main

import (
	"fmt"
)

//测试 slice

func show(info []int) {
	fmt.Printf("addr = %p \n", info)
	fmt.Printf("addr len = %v \n", len(info))
	fmt.Printf("addr cap = %v \n", cap(info))
	fmt.Println()
}

func main() {

	array1 := []int{1, 2, 3, 4, 5, 6}
	show(array1)
	show(array1[:0])

	fmt.Println("------------------------")

	array2 := []int{1, 2, 3, 4, 5, 6}
	show(array2)
	show(array2[0:0])

	fmt.Println("------------------------")

	array3 := []int{1, 2, 3, 4, 5, 6}
	show(array3)
	array3 = nil
	show(array3[0:0])

}
