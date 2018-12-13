package main

import (
	"fmt"
	"strconv"
)

//内部类型与外部类型实现相同的接口

//User user
type User struct {
	name     string
	birthday string
}

//Admin admin
type Admin struct {
	User
	level int
}

type inter interface {
	notify() string
}

func (u *User) notify() string {
	if u == nil {
		return ""
	}
	return "User:[" + u.name + ":" + u.birthday + "]"
}

func (a *Admin) notify() string {
	if a == nil {
		return ""
	}
	return "Admin:[" + a.name + ":" + a.birthday + ":" + strconv.Itoa(a.level) + "]"
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

	sendNotify(&admin)
	sendNotify(&admin.User)
	fmt.Printf("%v \n", admin.User.notify())
	fmt.Printf("%v \n", admin.notify())

}
