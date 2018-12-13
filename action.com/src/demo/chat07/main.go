package main

import (
	"fmt"
)

//嵌入式类型应用与接口

//User 用户
type User struct {
	name     string
	birthday string
}

//Admin 管理员
type Admin struct {
	User  //嵌入式类型
	level int
}

type inter interface {
	notify() string
}

func (u *User) notify() string {
	if u == nil {
		return ""
	}
	return "User[" + u.name + ":" + u.birthday + "]"
}

func sendNotify(it inter) {
	fmt.Printf("message: %v \n", it.notify())
}

func main() {

	admin := Admin{
		User: User{
			name:     "yzh",
			birthday: "1996-12-24",
		},
		level: 1,
	}

	sendNotify(&admin.User)
	sendNotify(&admin) //内部类型提升到外部类型

}
