package main

import "fmt"

//函数式编程

type iAdder func(int) (int, iAdder)

func adder(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder(base + v)
	}
}

func main() {

	add := adder(0)
	// fmt.Printf("%T %v \n", add, add)

	result, add := add(1)
	fmt.Printf("result = %d \n", result)
	result, add = add(2)
	fmt.Printf("result = %d \n", result)
	result, add = add(3)
	fmt.Printf("result = %d \n", result)
	result, add = add(4)
	fmt.Printf("result = %d \n", result)

}
