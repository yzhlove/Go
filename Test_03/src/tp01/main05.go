package main

import "strconv"

// 特殊"_"的使用

func main() {

	x,_ := strconv.Atoi("123")
	println("x = ",x)

}
