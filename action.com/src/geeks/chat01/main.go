package main

import (
	"fmt"
	"sync"
	"unsafe"
)

//只运行一次

type Singleton struct{}

var (
	singletonInstance *Singleton
	once              sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("create singleton instance .")
		singletonInstance = &Singleton{}
	})
	return singletonInstance
}

func main() {

	var wg sync.WaitGroup
	index := 10
	wg.Add(index)
	for i := 0; i < index; i++ {
		go func() {
			defer wg.Done()
			temp := GetInstance()
			fmt.Printf("%T %#X \n", temp, unsafe.Pointer(temp))
		}()
	}
	wg.Wait()
}
