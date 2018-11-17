package main

import "fmt"

// 切片操作

func main() {

	s := []int{0,1,2,3,4}

	p := &s
	p0 := &s[0]
	p1 := &s[1]

	println(p,p0,p1)

	(*p)[0] += 100
	*p1 += 100

	fmt.Println(s)

	(*p)[2] += 200
	fmt.Println(s)
}
