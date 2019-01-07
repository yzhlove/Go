package main

import "fmt"

//Go 1.9新功能 类型别名
//类型别名不允许定义方法

//类型定义
type NewInt int

//类型别名
type IntAsia = int

//Error 类型别名不允许用于方法
//func (i IntAsia) add(a,b int) {}

func main() {

	var a NewInt
	var b IntAsia

	fmt.Printf("a =  %T \n", a)
	fmt.Printf("b = %T \n", b)

}
