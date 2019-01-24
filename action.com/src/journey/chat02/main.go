package main

import (
	"fmt"
	"time"
)

//一个类型对应多个接口

type Info interface {
	show(data interface{})
}

type Detail interface {
	specific(data interface{})
}

type User struct {
	Name     string
	Age      uint
	Birthday string
}

func (u *User) show(data interface{}) {
	fmt.Println("show:", u)
}

func (u *User) specific(data interface{}) {
	fmt.Println("specific:", data)
}

func main() {
	user := new(User)
	user.Name = "xjj"
	user.Age = 18
	user.Birthday = time.Now().Format("2006/01/02")
	user.show(nil)

	user.specific(map[string]string{
		"name": "love xjj",
	})

}
