package main

import (
	"sync"

	"fmt"
)

//pool use

func main() {

	var counter int
	pools := sync.Pool{
		New: func() interface{} {
			counter += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	pools.Put(pools.New())
	pools.Put(pools.New())
	pools.Put(pools.New())
	pools.Put(pools.New())

	const workers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := workers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := pools.Get().(*[]byte)
			defer pools.Put(mem)
			//一些操作
		}()
	}
	wg.Wait()
	fmt.Printf("counter = %v \n", counter)

}
