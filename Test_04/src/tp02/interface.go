// package main

// import (
// 	"fmt"
// )

// //接口的使用

// //Animal 动物接口
// type Animal interface {
// 	Speak() string
// }

// type cat struct{}

// func (cat) Speak() string {
// 	return "cat speak!"
// }

// type dog struct{}

// func (dog) Speak() string {
// 	return "dag speak!"
// }

// type mouse struct{}

// func (mouse) Speak() string {
// 	return "mouse speak!"
// }

// type java struct{}

// func (java) Speak() string {
// 	return "java programmer!"
// }

// func main() {

// 	objs := []Animal{cat{}, dog{}, mouse{}, java{}}

// 	for _, aml := range objs {
// 		fmt.Println(aml.Speak())
// 	}

// }
