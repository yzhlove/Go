package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int
	var b string

	typeOfA := reflect.TypeOf(a)
	typeOfB := reflect.TypeOf(b)

	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(typeOfB.Name(), typeOfB.Kind())

}
