package main

import (
	"fmt"
	"strconv"
)

//接口值类型

//TypeInfo 显示类型
type TypeInfo interface {
	Show() string
}

func show(inter TypeInfo) {
	fmt.Printf("%v", inter.Show())
}

//固定类型
func showUser(inter interface{}) {
	fmt.Printf("%v ", inter.(*User).Show())
}

//User 学生
type User struct {
	Name string
	Age  int
}

//Teacher 老师
type Teacher struct {
	Name    string
	Project string
}

//Show User->Show方法
func (u *User) Show() string {
	return "Student: [" + u.Name + ":" + strconv.Itoa(u.Age) + "] \n"
}

//Show Teacher->Show方法
func (t *Teacher) Show() string {
	return "Teacher: [" + t.Name + ":" + t.Project + "] \n"
}

func main() {

	user := &User{
		Name: "jay",
		Age:  40,
	}

	teacher := &Teacher{
		Name:    "cdm",
		Project: "Language Class",
	}

	fmt.Printf("%v", user.Show())
	fmt.Printf("%v", teacher.Show())

	fmt.Println("--------------------")

	show(user)
	show(teacher)

	fmt.Println("--------------------")

	showUser(user)
	// showUser(teacher)

}
