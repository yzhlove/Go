package main

import (
	"fmt"
	"math/rand"
)

func A() {
	fmt.Println("A")
	B()
}

func B() {
	fmt.Println("B")
}

func C(callback func()) {
	fmt.Println("C")
	callback()
}

func D() {
	fmt.Println("D")
}

func E(value int) {
	fmt.Println("E = ", value)
}

func F() int {
	value := rand.Intn(100)
	fmt.Println("F = ", value)
	return value
}

func main() {

	fmt.Println("begin")
	defer A()
	fmt.Println("after")
	defer C(D)
	fmt.Println("first")
	defer E(F())
	fmt.Println("last")

}
