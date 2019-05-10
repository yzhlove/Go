package main

import "expression/chat04/obj"

func main() {

	obj.Init()

	//var basic obj.User
	//basic = &obj.Student{BaseInfo: new(obj.BaseInfo), StudentID: 123}
	//
	//basic.InitValue(&obj.BaseInfo{Name: "xjj", Age: 21, Birthday: "1996-05-23"})
	//
	//basic.UserInfo()

	stu := obj.UserList["STU"]
	stu.InitValue(&obj.BaseInfo{Name: "xjj", Age: 21, Birthday: "1996-05-23"})
	stu.UserInfo()

}
