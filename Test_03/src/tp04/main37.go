package main

import "fmt"

type User struct {
	Name    string
	Student []struct {
		Name string
		Age  int
	}
}

func main() {

	user := User{
		Name: "xdy",
	}

	user.Student = append(user.Student, struct {
		Name string
		Age  int
	}{Name: "xdy", Age: 28})

	user.Student = append(user.Student, struct {
		Name string
		Age  int
	}{Name: "yzh", Age: 18})

	fmt.Printf("%v \n", user)

}
