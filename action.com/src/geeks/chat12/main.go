package main

import (
	"fmt"
	"reflect"
)

//一个简单的反射例子

type Employee struct {
	EmployeeID uint64
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(value int) {
	e.Age = value
}

func main() {
	employee := &Employee{1234, "yurisa", 16}
	fmt.Println(reflect.ValueOf(*employee).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*employee).FieldByName("Name"); !ok {
		fmt.Println("name is not found ")
	} else {
		fmt.Println("Tag => ", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(employee).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(100)})

	fmt.Println(*employee)
}
