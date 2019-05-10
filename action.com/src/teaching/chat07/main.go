package main

import (
	"teaching/chat07/base"
	_ "teaching/chat07/student"
	_ "teaching/chat07/teacher"
)

func main() {

	stu := base.NewUser("student")
	stu.Info()

	teacher := base.NewUser("teacher")
	teacher.Info()

}
