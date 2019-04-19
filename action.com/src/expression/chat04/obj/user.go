package obj

import "fmt"

type User interface {
	UserInfo()
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
