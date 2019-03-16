package main

import "reflect"

func main() {

	var a int = 1024

	typeofA := reflect.ValueOf(a)

	typeofA.SetInt(100)

}
