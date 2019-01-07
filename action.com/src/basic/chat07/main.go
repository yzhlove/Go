package main

import "fmt"

//类型别名与结构体

type User struct {
	Name string
	Age  int
}

type Student = User

func (s Student) Show() {
	fmt.Printf("name = %s age = %d \n", s.Name, s.Age)
}

func main() {

	st := Student{
		Name: "lcm",
		Age:  22,
	}

	st.Show()

}
