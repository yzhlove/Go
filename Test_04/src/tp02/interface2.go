package main

import (
	"fmt"
)

//在接口转换的规则中，指针类型可以直接访问值类型，反之则不可行

//Animal 接口
type Animal interface {
	Speak() string
}

type cat struct{}

func (*cat) Speak() string {
	return "cat speak!"
}

type dog struct{}

func (dog) Speak() string {
	return "dog speak!"
}

func main() {

	animals := []Animal{new(cat), dog{}}

	for _, it := range animals {
		fmt.Println(it.Speak())
	}

}
