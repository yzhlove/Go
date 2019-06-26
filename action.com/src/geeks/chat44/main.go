package main

import "fmt"

//逃逸分析

type User struct {
	Name string
	Age  int
}

func GetUser(user *User) *User {
	return user
}

func main() {

	var u *User

	func() {
		u = GetUser(&User{Name: "yzh", Age: 20})
	}()

	fmt.Printf("%p ", u)

}
