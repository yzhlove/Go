package main

import "fmt"

//函数测试

type class struct{}

func (slef class) Do(value int) int {
	fmt.Println("class is function:", value)
	return 0
}

func funcDo(value int) int {
	fmt.Println("standard value:", value)
	return 0
}

func main() {

	var tf func(int) int

	var c class
	tf = c.Do

	tf(100)

	tf = funcDo
	tf(200)

}
