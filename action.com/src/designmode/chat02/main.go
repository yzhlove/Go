package main

import "fmt"

//测试

type User struct {
	Name     string
	Age      int
	Birthday string
}

type Project interface {
	InfoShow()
}

func (u *User) Show() {
	fmt.Printf("User:%v \n", u)
}

type Student struct {
	User
	StudentID int
}

func (s *Student) InfoShow() {
	fmt.Printf("[Student => %v]  \n", s)
}

type Teacher struct {
	User
	TeacherID int
}

func (t *Teacher) InfoShow() {
	fmt.Printf("[Teacher => %v] \n", t)
}

func main() {

	s := &Student{User{Name: "yzh", Age: 20, Birthday: "1996-12-24"}, 100123}
	s.Show()
	s.InfoShow()

	t := &Teacher{User{Name: "xjj", Age: 21, Birthday: "1996-05-23"}, 700123}
	t.Show()
	t.InfoShow()

}
