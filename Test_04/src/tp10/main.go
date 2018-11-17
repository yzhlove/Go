package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGrop应用

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("go run 1")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("go run 2")
	}()

	go func() {
		fmt.Println("go run 3")
	}()

	wg.Wait()

}
