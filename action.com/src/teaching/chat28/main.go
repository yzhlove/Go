package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 1024
	valueOfA := reflect.ValueOf(a)

	var getA int = valueOfA.Interface().(int)

	var getB int64 = valueOfA.Int()

	fmt.Println(getA, " - ", getB)

}
