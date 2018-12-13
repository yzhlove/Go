package main

import (
	"demo/chat11/examp"
	"fmt"
)

//嵌入类型的字段初始化的技巧

func main() {

	//创建管理员

	admin := examp.Admin{
		Rights: 100,
	}
	//初始化user
	admin.Name = "yzh"
	admin.Email = "lcm5201314@gmail.com"

	fmt.Printf("admin:%v \n", admin)

}
