package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a *int
	fmt.Println("a *int:", reflect.ValueOf(a).Type(), reflect.ValueOf(a).IsNil())

	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	s := struct{}{}

	fmt.Println("no member:", reflect.ValueOf(s).FieldByName("").IsValid())

	fmt.Println("no function", reflect.ValueOf(s).MethodByName("").IsValid())

	m := map[int]int{}
	fmt.Println("no key:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())

}
