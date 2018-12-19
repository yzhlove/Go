package main

import "fmt"

//foeach 结构体

//User 用户
type User struct {
	name     string
	birthday string
	sex      int
}

func main() {

	user := &User{
		name:     "yzh",
		birthday: "1996-12-24",
		sex:      18,
	}

	fmt.Printf("User = %v \n", user)

}
