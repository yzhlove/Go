package main

import "fmt"

//private 与 public 的应用

//InitNumber 新的类型
type InitNumber int

func main() {

	//新的类型的初始化的方式
	var it InitNumber = 100
	fmt.Printf("%v \n", it)

	fmt.Printf("%v \n", InitNumber(200))

	itt := new(InitNumber)
	*itt = 12138
	fmt.Printf("%v \n", *itt)

}
