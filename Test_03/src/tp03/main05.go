package main

import "fmt"

func testb(a ...int) {
	fmt.Printf("%T,%v\n",a,a)
}

func main() {

	a := [3]int {10,20,30}
	testb(a[:]...) 	// 转化为切片后展开

}