package main

import "fmt"

// 自定义类型

func main() {

	type (
		user struct {				// 结构体
			name string
			age uint8
		}
		event func(string) bool		// 函数类型
	)

	var u user = user{"Tom",22}
	fmt.Println(u)

	admin := user{"Admin",19}
	fmt.Printf("%T,%v",admin,admin)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")


}
