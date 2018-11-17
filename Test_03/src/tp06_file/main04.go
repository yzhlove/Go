package main

import (
	"sync"
	"time"
)

//Mutex的应用
func main() {

	var lock sync.Mutex

	m := make(map[string]int)

	go func() {
		for {
			lock.Lock()
			m["a"] = 1
			lock.Unlock()
		}

	}()

	go func() {

		for {
			lock.Lock()
			_ = m["b"]
			time.Sleep(time.Microsecond)
			lock.Unlock()
		}

	}()

	select {}

}
