package main

import (
	"fmt"
	"reflect"
)

type Cat struct{}

func main() {

	ins := &Cat{}

	typeOfCat := reflect.TypeOf(ins)
	fmt.Println("name = ", typeOfCat.Name(), " kind = ", typeOfCat.Kind())

	typeOfcat := typeOfCat.Elem()
	fmt.Println("element name = ", typeOfcat.Name(), "element kind = ", typeOfcat.Kind())

}
