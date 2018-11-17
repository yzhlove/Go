package main

import (
	"fmt"
)

//空借口类型准换
func main() {

	var x interface{} = func(n int) string {
		return fmt.Sprintf("data:%d ", n)
	}

	switch v := x.(type) {
	case nil:
		fmt.Println("nil")
	case func(int) string:
		fmt.Println(v(1000))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		fmt.Println("unkonow")
	}

}
