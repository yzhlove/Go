package main

import (
	"fmt"
)

//go 多态

//User 用户
type User struct {
	name     string
	birthday string
}

//Admin  管理员
type Admin struct {
	name     string
	password string
	email    string
}

type notiter interface {
	notify() string
}

func sendMessage(not notiter) {
	fmt.Printf("message: %v \n", not.notify())
}

//实现接口

//notify 接口
func (u *User) notify() string {
	if u != nil {
		return "User[" + u.name + ":" + u.birthday + "]" + "\n"
	}
	return ""
}

func (a *Admin) notify() string {
	if a != nil {
		return "Admin[" + a.name + ":" + a.email + "]" + "\n"
	}
	return ""
}

func main() {

	u := User{
		name:     "yzh",
		birthday: "1996-12-24",
	}

	a := Admin{
		name:     "xjj",
		password: "******",
		email:    "xjjlove@gmail.com",
	}

	sendMessage(&u)
	sendMessage(&a)

}
