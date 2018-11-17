package main

import (
	"fmt"
)

type file struct {
	name string
}

type datas struct {
	file
	name string
}

func main() {

	//重复命名

	d := datas{
		name: "fist_name",
		file: file{name: "last_name"},
	}

	fmt.Printf("%+v \n", d)

	d.name = "what are you doing"
	d.file.name = "where are you going"
	fmt.Printf("%+v \n", d)

}
