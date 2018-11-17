package main

import (
	"fmt"
)

//方法调用

//N 类型
type N int

func (n N) test() {
	fmt.Printf("== %+v \n", n)
}

func main() {

	var n N
	n = 100
	fn := N.test
	fn(n)

}
