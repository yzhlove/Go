package main

import "fmt"

//斐波拉切数列

func fbi() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {

	fb := fbi()
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())
	fmt.Printf("fd : %v \n", fb())

}
