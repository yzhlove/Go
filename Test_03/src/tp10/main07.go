package main

import (
	"fmt"
)

//方法覆盖

type user struct {
	name string
	age  int
}

type manager struct {
	user
	project string
}

func (u *user) toString() {
	fmt.Printf("user: %+v \n", *u)
}

func (m *manager) toString() {
	m.user.toString()
	fmt.Printf("manager: %+v \n", m)
}

func main() {

	var m manager
	m.user.name = "kaka"
	m.user.age = 18
	m.project = "java"

	m.toString()

}
