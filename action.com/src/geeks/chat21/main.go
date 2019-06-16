package main

import (
	"sync"

	"fmt"
)

//once 单例模式

var counter int

func main() {

	var (
		once      sync.Once
		increFunc = func() { counter++ }
		wg        sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(increFunc)
		}()
	}

	wg.Wait()

	fmt.Println("counter = ", counter) //1

}
