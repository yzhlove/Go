package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func a() {
	mutex.Lock()
	fmt.Println("a")
	mutex.Unlock()
}

func b() {
	mutex.Lock()
	fmt.Println("b")
	a()
	mutex.Unlock()
}

func main() {
	b()
}
