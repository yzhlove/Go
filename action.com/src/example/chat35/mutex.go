package main

import (
	"fmt"
	"sync"
)

//重复加锁

var mutex sync.Mutex

func add() {
	mutex.Lock()
	fmt.Println(" add Mutex")
	mutex.Unlock()
}

func main() {

	var addMutex sync.Mutex

	addMutex.Lock()
	add()
	addMutex.Unlock()

}
