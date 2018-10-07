package main

import "fmt"

// 指针传递与值传递

func test(x [2]int) {
	fmt.Printf("x : %p ,%v \n",&x,x)
}

func test2(x *[2]int) {
	fmt.Printf("y : %p ,%v \n",&x,x)
}

func main() {

	a := [2]int{10,20}
	var b [2]int
	b = a
	fmt.Printf("a:%p , %v \n",&a,a)
	fmt.Printf("b:%p , %v \n",&b,b)

	test(a)
	test2(&a)

}

/*
a:0xc000078010 , [10 20]
b:0xc000078020 , [10 20]
x : 0xc000078060 ,[10 20]
y : 0xc000088020 ,&[10 20]
*/
