package main

import "fmt"

// 字典

// 创建

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct{x int}{
		1:{x:100},
		2:{x:200},
	}

	fmt.Println(m,m2)

}
