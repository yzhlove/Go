package main

import "fmt"

func testc(a ...int) {

	for i := range a {
		a[i] += 100
	}

}


func main() {

	// 数组
	a := []int {10,20,30}
	testc(a...)
	fmt.Println(a)

	// 切片
	b := [3]int {10,20,30}
	testc(b[:]...)
	fmt.Println(b)
}
