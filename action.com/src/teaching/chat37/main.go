package main

import "fmt"

type Add func(int, int) int

func (*Add) Info() {
	fmt.Println("What are you doing")
}

func tmp(a, b int) int {
	return a + b
}

func Tmp(fn Add, a, b int) {
	fmt.Println(fn(a, b))
}

func main() {

	add := new(Add)
	add.Info()

	f := *add
	f.Info()

	Tmp(tmp, 10, 20)

}
