package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}


func live() People {
	var stu *Student
	return stu
}

func main() {
	lv := live()
	fmt.Printf("%T\n", lv)
	fmt.Printf("%T\n", nil)
	if lv == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
