package main

import "fmt"

// 自定义类型

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {

	f := read | exec
	fmt.Printf("%b \n",f)


}