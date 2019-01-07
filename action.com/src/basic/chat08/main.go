package main

import (
	"fmt"
	"reflect"
)

//类型别名与结构体

type Brand struct{}

func (t Brand) Show() {}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {

	var a Vehicle

	a.FakeBrand.Show()

	ta := reflect.TypeOf(a)
	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FiledName: %v FieldType:%v \n", f.Name, f.Type.Name())
	}

}
