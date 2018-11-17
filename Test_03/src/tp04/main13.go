package main

import "fmt"

// 数组的初始化方式

func main() {

	var a [4]int		// 元素自动初始化为0
	b := [4]int{2,5}	// 未提供初始指的元素自动初始化为0

	c := [4]int{5,3:10}	// 指定索引位置初始化
	d := [...]int{1,2,3}
	e := [...]int{10,3:100}	// 索引为3地方初始化为100

	fmt.Println(a,b,c,d,e)

}
