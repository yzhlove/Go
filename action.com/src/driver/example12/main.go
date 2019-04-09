package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u *User) ShowInfo() {
	fmt.Printf("%v\n", u)
}

func main() {
	u := new(User)
	u.ShowInfo()
	u.Name = "xjj"
	u.ShowInfo()
	u.Age = 20
	u.ShowInfo()
	u.Age = 21
	u.ShowInfo()
}
