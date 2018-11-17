package main

import (
	"fmt"
)

//指针使用
func main() {

	type user struct {
		name string
		age  int
	}

	p := &user{
		name: "Tom",
		age:  20,
	}

	fmt.Printf("%+v\n", *p)

	p.name = "Jerry"
	p.age++

	fmt.Printf("%+v\n", *p)

	u := p

	fmt.Printf("%T,%+v \n", u, u)

	u.name = "Uname"
	u.age++

	fmt.Printf("%+v \n", u)

	u2 := &p
	(*u2).name = "wht are you doing"
	(*u2).age++

	fmt.Printf("%+v \n", *u2)

}
