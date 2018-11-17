package main

import (
	"fmt"
)

//携程初步

//FunString 函数
type FunString func(interface{}) string

func show(i interface{}) {
	for _, v := range i.([]int) {
		fmt.Printf("%+v \n", v)
	}
}

func main() {

	tmp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fn := FunString(func())
	show(tmp)

}
