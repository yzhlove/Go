package main

import (
	"fmt"
	"sync"
)

//对象池

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(200)
	pool.Put(300)

	index := 10
	var wg sync.WaitGroup
	wg.Add(index)
	for i := 0; i < index; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("Get Pool => %v \n", pool.Get().(int))
		}()

	}
	wg.Wait()

}
