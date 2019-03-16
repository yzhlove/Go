package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 1024
	A := reflect.ValueOf(&a)

	A.Elem().SetInt(2048)

	fmt.Println("value = ", A.Elem().Int())

}
