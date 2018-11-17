package main

import (
	"fmt"
	"runtime"
)

//Go让出时间片

func say(s string) {
	for i := 0; i <= 2; i++ {
		fmt.Println("s = ", s)
		runtime.Gosched()
	}
}

func main() {

	go say("world")
	say("hello")
}
