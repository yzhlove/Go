package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "fly", "skill is action")

func main() {

	flag.Parse()

	var funcList = map[string]func(){
		"fire": func() {
			fmt.Println("FIRE")
		},
		"run": func() {
			fmt.Println("RUN")
		},
		"fly": func() {
			fmt.Println("FLY")
		},
	}

	if fn, ok := funcList[*skillParam]; ok {
		fn()
	} else {
		fmt.Println("Not Found Func!")
	}

}
