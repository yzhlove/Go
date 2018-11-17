package main

import "fmt"

// 第三章

// 隐式类型转换

func main() {

	const v = 20
	var a byte = 10

	b := v + a
	fmt.Printf("%T,%v",b,b)

	const c float32 = 1
	d := c + v
	fmt.Printf("%T,%v",d,d)

}




