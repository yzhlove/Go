package main

import "fmt"

//闭包的记忆效应

func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		value++
		return value
	}
}

func main() {
	//创建一个累加器
	accmulate := Accumulate(1)
	fmt.Println(accmulate())
	fmt.Println(accmulate())

	fmt.Printf("%p \n", &accmulate)

	accmulate2 := Accumulate(10)
	fmt.Println(accmulate2())
	fmt.Println(accmulate2())
	fmt.Println(accmulate())
	fmt.Println(accmulate())
	fmt.Printf("%p \n", &accmulate2)

}
