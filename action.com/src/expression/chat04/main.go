package main

import "expression/chat04/obj"

func main() {

	var basic obj.User
	basic = &obj.Student{BaseInfo: new(obj.BaseInfo), StudentID: 123}

	basic.InitValue(&obj.BaseInfo{Name: "xjj", Age: 21, Birthday: "1996-05-23"})

	basic.UserInfo()

}
