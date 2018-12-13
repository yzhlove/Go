package main

import (
	"fmt"
)

//值与指针

type duration int

func (t *duration) pretty() {
	fmt.Printf("%T %v %v \n", t, t, *t)
}

//如果实现接口的

func main() {

	// duration(42).pretty()

}
