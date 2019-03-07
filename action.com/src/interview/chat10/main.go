package main

import (
	"fmt"
	"sync"
)

func main() {

	td := &thread{
		data: []string{"hello", "world"},
	}

	read := td.Iter()

	for {
		if elem, ok := <-read; ok {
			fmt.Println("element = ", elem)
		} else {
			break
		}
	}

}

type thread struct {
	sync.RWMutex
	data []string
}

func (set *thread) Iter() <-chan string {
	out := make(chan string)
	go func() {
		set.RLock()
		for _, elem := range set.data {
			out <- elem
		}
		close(out)
		set.RUnlock()
	}()
	return out
}
