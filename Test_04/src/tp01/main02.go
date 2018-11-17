package main

import (
	"fmt"
	"strconv"
)

//接口包含 [类似继承]
type tester interface {
	test()
}

type stringer interface {
	tester
	ToString() string
}

type user struct {
	name string
	age  int
}

func (u *user) test() {
	fmt.Println("test a b c")
}

func (u *user) ToString() string {
	return u.name + " :: " + strconv.Itoa(u.age)
}

func show(s stringer) {
	s.test()
	fmt.Printf("%#v \n", s.ToString())
}

func main() {

	u := &user{
		name: "yzh", age: 18,
	}
	show(u)
}
