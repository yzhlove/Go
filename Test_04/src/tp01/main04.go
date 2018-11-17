package main

import (
	"fmt"
)

//接口复制的是值,要在接口里面改变对象的直，则徐复制指针

type student struct {
	project string
	name    string
	age     int
}

func main() {

	st := student{
		project: "English",
		name:    "yzh",
		age:     18,
	}

	var d1 interface{} = st
	var d2 interface{} = &st

	fmt.Printf("%+v \n", st)
	fmt.Printf("%+v \n", d1.(student))
	fmt.Printf("%+v \n", d2.(*student))

	// d1.(student).project = "math"	//值复制，无法被修改
	fmt.Printf("%+v \n", st)
	d2.(*student).name = "xjj"
	fmt.Printf("%+v \n", st)

}
