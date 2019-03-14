package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mutex sync.RWMutex
)

func GetCount() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return count
}

func SetCount(value int) {
	mutex.Lock()
	count = value
	mutex.Unlock()
}

func main() {

	SetCount(100)
	fmt.Println(GetCount())

}
