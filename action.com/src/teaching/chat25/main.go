package main

import (
	"fmt"
	"reflect"
)

type Enum int

const Zero Enum = 0

type Cat struct{}

func main() {

	typeOfCat := reflect.TypeOf(Cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	typeOfA := reflect.TypeOf(Zero)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

}
