package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2

	fmt.Printf("A : a = %d b = %d \n", a, b)
	defer calc("1", a, calc("10", a, b))
	fmt.Printf("B : a = %d b = %d \n", a, b)
	a = 0

	defer calc("2", a, calc("20", a, b))
	fmt.Printf("C : a = %d b = %d \n", a, b)
	b = 1

}
