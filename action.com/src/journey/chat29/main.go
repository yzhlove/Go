package main

import "fmt"

//接口赋值

func main() {

	var inter interface{}
	inter = 10
	fmt.Printf("%T %+v %d \n", inter, inter, inter)

	inter = "Hello World"
	fmt.Printf("%T %+v %s \n", inter, inter, inter)

	inter = false
	fmt.Printf("%T %+v %v \n", inter, inter, inter)

	//接口赋值
	var a = inter.(bool)
	fmt.Printf("a = %v \n", a)

}
