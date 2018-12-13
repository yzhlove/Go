package main

import (
	"fmt"
)

type sstr string

func (str *sstr) changeString() string {
	*str = "Hello World"
	return "what are you doing"
}

func main() {

	var s sstr
	stemp := s.changeString()
	fmt.Printf("stemp = %v s = %v \n", stemp, s)

}
