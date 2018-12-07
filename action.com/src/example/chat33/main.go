package main

import (
	"fmt"
	"reflect"
)

type hello struct {
	text string
}

type world struct {
	text string
}

type msg struct {
	msgHead interface{}
	msg     string
}

func comp(a interface{}, b map[string]interface{}) {
	for _, st := range b {
		fmt.Printf("%v %v\n", reflect.TypeOf(a), reflect.TypeOf(st))

		if reflect.TypeOf(a) == reflect.TypeOf(st) {
			fmt.Printf("yes : %v \n", st)
		}
	}
}

func main() {

	he := &hello{text: "Hello"}
	wr := &world{text: "World"}
	mapList := make(map[string]interface{})
	mapList["helloType"] = he
	mapList["worldType"] = wr
	comp(he, mapList)

}
