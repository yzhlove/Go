package main

import "fmt"

//  new使用

func main() {

	p := new(map[string]int)
	m := *p
	m["a"] = 1	// 错误，new仅仅只分配了指针对象的内存
	fmt.Println(m)

}