package obj

import (
	"fmt"
)

var UserList = make(map[string]User)

type User interface {
	UserInfo()
	Value
}

type Value interface {
	InitValue(info *BaseInfo)
}

type BaseInfo struct {
	Name     string
	Age      int
	Birthday string
}

func (info *BaseInfo) InitValue(temp *BaseInfo) {
	info.Name = temp.Name
	info.Age = temp.Age
	info.Birthday = temp.Birthday
}

type Student struct {
	*BaseInfo
	StudentID int
}

func (student *Student) UserInfo() {
	fmt.Println("Student Info : ", student, student.Name, student.Age, student.Birthday, student.StudentID)
}

type Teacher struct {
	*BaseInfo
	TeacherID int
}

func (teacher *Teacher) UserInfo() {
	fmt.Println("Teacher Info : ", teacher)
}

func Init() {
	UserList["STU"] = &Student{BaseInfo: new(BaseInfo)}
	UserList["TEH"] = &Teacher{BaseInfo: new(BaseInfo)}
}
