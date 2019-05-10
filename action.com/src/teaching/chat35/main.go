package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int
	typeOfA := reflect.TypeOf(a)

	aIns := reflect.New(typeOfA)
	fmt.Println(aIns.Type(), aIns.Kind())

}
