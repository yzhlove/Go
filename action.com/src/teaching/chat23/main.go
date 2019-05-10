package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//RWMutex测试

var (
	mutexValue = 1
	mutex      sync.RWMutex
)

func Set(value int) {
	mutex.Lock()
	mutexValue = value
	fmt.Println(" ------> set value :", value)
	time.Sleep(time.Second)
	mutex.Unlock()
}

func Get(id int) {
	mutex.RLock()
	fmt.Printf("[ID: %d value:%d ] \n", id, mutexValue)
	time.Sleep(time.Millisecond * 500)
	mutex.RUnlock()
}

func main() {

	for i := 10; i < 100; i++ {
		go Get(i)
	}
	for i := 0; i < 3; i++ {
		go Set(rand.Intn(100) + 50)
	}
	for i := 10; i < 100; i++ {
		go Get(i)
	}

	time.Sleep(time.Second * 20)

}
