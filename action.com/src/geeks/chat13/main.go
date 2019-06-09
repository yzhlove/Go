package main

import (
	"fmt"
	"reflect"
)

//map 和 slice 只能和 nil 比较
//但是可以利用放射的方式比较内容

func main() {

	a := map[int]string{1: "i", 2: "love", 3: "you"}
	b := map[int]string{1: "i", 2: "love", 3: "you"}

	//fmt.Println(a == b)	//Error:map can only be compared to nil
	fmt.Println(reflect.DeepEqual(a, b)) //true

	c := []int{1, 2, 3, 4, 5}
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.DeepEqual(c, d)) //true
}
