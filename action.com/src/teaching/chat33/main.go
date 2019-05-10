package main

import (
	"fmt"
	"reflect"
)

func main() {

	type Dog struct {
		age int
	}

	dog := reflect.ValueOf(Dog{})
	ageref := dog.FieldByName("age")
	fmt.Println("age = ", ageref.Int())
	ageref.SetInt(1234)

}
