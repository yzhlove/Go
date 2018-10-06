package main

import (
	"github.com/pkg/errors"
	"log"
)

// 错误处理

var errDivByZero = errors.New("divsion by zero")

func div(x,y int) (int,error) {
	if y == 0 {
		return 0,errDivByZero
	}
	return x/y ,nil
}


func main() {
	z,err  := div(5,0)
	if err == errDivByZero {
		log.Fatalln(err)
	}
	println(z)
}