package main

import (
	"fmt"
)

//匿名接口的使用

type data struct{}

func (data) sstring() string {
	return "1234"
}

type node struct {
	data interface {
		sstring() string
	}
}

func main() {

	var t interface {
		sstring() string
	} = data{}

	n := node{
		data: t,
	}

	fmt.Println(n.data.sstring())

}
