package main

import (
	"fmt"
	"reflect"
)

func main() {
	// failure
	//type Dog struct {
	//	Age int
	//}
	//refAge := reflect.ValueOf(Dog{})
	//refValue := refAge.FieldByName("Age")
	//fmt.Println("value = ", refValue.Int())
	//refValue.SetInt(1234)

	type Dog struct {
		Age int
	}

	refAge := reflect.ValueOf(&Dog{})
	refAgeValue := refAge.Elem()
	ref := refAgeValue.FieldByName("Age")
	fmt.Println("refValue = ", ref.Int())

	ref.SetInt(1024)
	fmt.Println("refValue = ", ref.Int())
}
