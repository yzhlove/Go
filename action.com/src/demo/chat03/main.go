package main

import (
	"fmt"
)

//接口类型

type interier interface {
	notify()
}

//User 用户
type User struct {
	name string
	sex  int
}

func (u *User) notify() {
	fmt.Printf("%T %v \n", u, u)
}

func main() {

	u := &User{
		name: "yzh",
		sex:  1,
	}
	showNotify(u)
}

func showNotify(it interier) {
	it.notify()
}
