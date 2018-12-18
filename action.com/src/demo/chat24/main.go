package main

import "fmt"

//闭包与自由变量

//闭包
func addFunction() func(int) int {
	sum := 0 //自由变量
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {

	addfunc := addFunction()
	for i := 0; i <= 10; i++ {
		fmt.Printf("addfunc(%d)  = %d \n", i, addfunc(i))
	}
}

/*
addfunc(0)  = 0
addfunc(1)  = 1
addfunc(2)  = 3
addfunc(3)  = 6
addfunc(4)  = 10
addfunc(5)  = 15
addfunc(6)  = 21
addfunc(7)  = 28
addfunc(8)  = 36
addfunc(9)  = 45
addfunc(10)  = 55
*/
