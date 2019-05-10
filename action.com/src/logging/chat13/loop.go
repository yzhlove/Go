package main

import "time"

var done bool

func setup() {
	time.Sleep(time.Second)
	done = true
}

func main() {
	go setup()
	for !done {
	}
}
