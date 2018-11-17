package main

import (
	"fmt"
)

//多级指针调用

//User 用户信息
type User struct {
	name string
	age  int
}

func (u User) reloadValue() {
	u.name = "reload_value"
	u.age = 100
	fmt.Printf("%#v \n", u)
}

func (u *User) reloadPoint() {
	u.name = "reload_point"
	u.age = 100
	fmt.Printf("%#v \n", *u)
}

func main() {

	u := User{}

	pu := &u

	fmt.Printf("global : %+v \n", u)

	u.reloadValue()

	fmt.Printf("global : %+v \n", u)

	u.reloadPoint()

	fmt.Printf("global : %+v \n", u)

	fmt.Printf("global : %+v | %+v \n", pu, *pu)

	fmt.Println("-----------------------------")

	pu.name = "point_value"
	pu.age = 12344
	fmt.Printf("global : %+v \n", pu)
	pu.reloadValue()
	fmt.Printf("global : %+v \n", pu)
	pu.reloadPoint()
	fmt.Printf("global : %+v \n", pu)

}
