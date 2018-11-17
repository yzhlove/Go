package main

import (
	"fmt"
)

//匿名结构体

func main() {

	user := struct {
		name string
		age  int
	}{
		name: "Tom",
		age:  21,
	}

	type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{
		name: "test.dat",
	}
	f.attr.owner = 1
	f.attr.perm = 0755

	fmt.Println(user, f)
}
