package main

import (
	"fmt"
)

//关于空结构题
func main() {

	var d [100]struct{}

	s := d[:]
	fmt.Println(s[3], len(d), cap(d))

	d[1] = struct{}{}
	d[2] = struct{}{}

	fmt.Println(s, d)

}
