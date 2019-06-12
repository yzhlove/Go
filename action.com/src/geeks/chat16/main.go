package main

import "fmt"

func main() {

	a := make([]uint8, 0, 8)
	fmt.Printf("%p - %v - %v \n", a, len(a), cap(a))
	a = append(a, 1, 2, 3, 4)
	b := a
	fmt.Printf("%p - %v - %v \n", a, len(a), cap(a))
	fmt.Printf("%p - %v - %v \n", b, len(b), cap(b))

	b = append(b, 5)
	fmt.Printf("%p - %v - %v \n", a, len(a), cap(a))
	fmt.Printf("%p - %v - %v \n", b, len(b), cap(b))

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	b[3] = 7
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
