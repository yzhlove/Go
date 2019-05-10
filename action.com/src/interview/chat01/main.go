package main

import (
	"errors"
	"fmt"
)

func main() {

	test()

}

func test() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	panic(errors.New("error !!!"))
}
