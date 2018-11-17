package main

import (
	"fmt"
	"strconv"
)

//接口的使用 实现接口
type tester interface {
	test()
	ToString() string
}

type data struct {
	name string
	age  int
}

func (d *data) test() {
	fmt.Println("test...")
}

func (d *data) ToString() string {
	return d.name + " : " + strconv.Itoa(d.age)
}

func main() {

	var d data
	d = data{name: "yzh", age: 16}
	var t tester = &d

	t.test()
	fmt.Println(t.ToString())

}
