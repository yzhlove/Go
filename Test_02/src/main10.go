package main

import "fmt"

type student struct {
	name string
	age int
}

func (stu student) Print()  {
	fmt.Printf("%+v\n",stu)
}


type Printer interface {
	Print()
}


func main() {

	var u student
	u.name  = "fuck"
	u.age = 16

	var inter Printer = u
	inter.Print()

}
