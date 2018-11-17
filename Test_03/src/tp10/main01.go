package main

import (
	"fmt"
)

type data struct {
	*int
	string
}

func main() {

	x := 100
	d := data{
		int:    &x,
		string: "abc",
	}

	fmt.Printf("%#v \n", d)

}
