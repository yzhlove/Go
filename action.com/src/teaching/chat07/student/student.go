package student

import (
	"fmt"
	"teaching/chat07/base"
)

type Student struct{}

func (stu *Student) Info() {
	fmt.Println("<- student ->")
}

func init() {
	base.Register("student", func() base.User {
		return new(Student)
	})
}
