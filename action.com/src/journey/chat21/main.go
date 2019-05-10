package main

import "fmt"

//测试defer

func testDefer() {
	defer fmt.Println("1")
	{
		defer fmt.Println("2")
		defer fmt.Println("3")
	}

	func() {
		defer fmt.Println("4")
		defer fmt.Println("5")
	}()

	defer fmt.Println("6")
}

func main() {
	testDefer()
}
