package main

import "fmt"

//  可变参数

func test(s string,a ... int) {
	fmt.Printf("%T,%v\n",a,a)
}

func main() {
	test("abc",1,2,3,4)
}