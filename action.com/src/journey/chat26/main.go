package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	user := User{
		Name: "yzh",
		Age:  18,
	}

	fmt.Println("length ", len(user.Name))

}
