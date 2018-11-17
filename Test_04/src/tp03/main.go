package main

import (
	"fmt"
)

//接口类型准换

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

func main() {

	var d data = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok {
		fmt.Println("stringer: ", n)
	}

	if d2, ok := x.(data); ok {
		fmr.Println(d2)
	}

	e := x.(error)
	fmt.Println(e)

}
