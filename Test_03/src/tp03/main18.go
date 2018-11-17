package main

import (
	"fmt"
	"log"
)

// 自定义错误类型

type DivError struct {
	x,y int
}

func (DivError) Error() string {
	return "div by zero"
}

func div(x,y int) (int ,error) {
	if y == 0 {
		return 0,DivError{x,y}
	}
	return x / y,nil
}


func main() {

	z ,err := div(5,0)

	if err != nil {
		switch e := err.(type) {
		case DivError:
			fmt.Println(e,e.x,e.y)
		default:
			fmt.Println(e)
		}
		log.Fatalln(err)
	}
	println(z)
}