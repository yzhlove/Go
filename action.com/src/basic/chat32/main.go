package main

import "fmt"

//闭包

func main() {

	str := "hello world"

	foo := func() {
		str = "hello dude"
	}

	foo()
	fmt.Println(str)
}
