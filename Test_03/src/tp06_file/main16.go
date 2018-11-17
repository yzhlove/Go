package main

import (
	"fmt"
)

//FuncString 定义函数
type FuncString func() string

func (f FuncString) string() string {
	return f()
}

func main() {

	var t fmt.Stringer = FuncString(func() string {
		return "Hello World!"
	})

	fmt.Printf(t)

}
