package main

import "fmt"

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Error : %v \n", err)
		} else {
			panic(r)
		}
	}()

	//b := 0
	//a := 5 / b
	//fmt.Printf("a = %v \n", a)

	panic(123)

}

func main() {

	tryRecover()

}
