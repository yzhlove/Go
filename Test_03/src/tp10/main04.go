package main

import (
	"fmt"
)

//N int
type N int

func (n N) value() {
	n++
	fmt.Println("n = ", n)
}

func (n *N) point() {
	*n++
	fmt.Println("n = ", *n)
}

func main() {

	var n N = 25
	fmt.Println("global n = ", n)

	n.value()
	fmt.Println("global n = ", n)

	n.point()
	fmt.Println("global n = ", n)

}
