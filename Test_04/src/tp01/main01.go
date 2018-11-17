package main

import (
	"fmt"
)

//接口比较

func main() {

	var t1, t2 interface{}

	fmt.Println(t1 == t2)

	t1, t2 = 100, 100

	fmt.Println(t1 == t2)

	t1, t2 = map[int]int{}, map[int]int{}

	fmt.Println(t1 == t2)

}
