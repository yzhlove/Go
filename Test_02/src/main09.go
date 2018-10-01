package main

// 继承

import "fmt"

type people struct{
	name string
	age int
	birthday string
}

type company struct {
	people
	title string
}

func (u people) ToString() string {
	println(u.name," -- ",u.age)
	return fmt.Sprintf("%+v",u)
}


func main() {
	var comp company
	comp.name = "Tom"
	comp.age = 22

	println(comp.ToString())

}