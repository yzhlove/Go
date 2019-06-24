package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (info *User) Init(user User) {
	*info = user
}

func main() {

	a := User{}
	b := User{Name: "xyj", Age: 23}

	a.Init(b)

	fmt.Println(a)

}
