package main

import (
	"fmt"
	"log"
)

//控对象不等于 nil

//TestError 类型
type TestError struct{}

func (*TestError) Error() string {
	return "<error>"
}

func test(x int) (int, error) {
	var err *TestError

	if x < 0 {
		err = new(TestError)
		x = 0
	} else {
		x += 100
	}
	return x, err
}

func main() {

	x, err := test(-1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v \n", x)
}
