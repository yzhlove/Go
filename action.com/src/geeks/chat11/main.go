package main

import (
	"fmt"
	"reflect"
)

//一个简单的反射例子

func main() {

	var value uint32 = 128

	fmt.Println(reflect.TypeOf(value))         //instance
	fmt.Println(reflect.ValueOf(value))        //value
	fmt.Println(reflect.ValueOf(value).Type()) //value -> instance

}
