package main

import (
	"fmt"
)

//嵌入式类型:内部类型与外部类型的联系

//User 用户
type User struct {
	name     string
	birthday string
}

//Admin 管理员
type Admin struct {
	User     //嵌入式类型
	level    int
	password string
	email    string
}

func (u *User) notify() string {
	if u == nil {
		return ""
	}
	return "User[" + u.name + ":" + u.birthday + "]"
}

func main() {

	admin := Admin{
		User: User{
			name:     "yzh",
			birthday: "1996-12-24",
		},
		level:    80,
		password: "123456",
		email:    "lcmm5201314@gmail.com",
	}

	fmt.Printf("adminInfo:%v \n", admin.User.notify())

	fmt.Printf("adminInfo:%v \n", admin.notify())

	fmt.Printf("Name:%v  birthday:%v email:%v \n", admin.User.name, admin.birthday, admin.email)
}
