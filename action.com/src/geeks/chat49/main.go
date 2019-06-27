package main

import (
	"fmt"
	"math/rand"
)

func main() {

	repeatedFn := func(done <-chan struct{}, fn func() interface{}) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			select {
			case stream <- fn():
			case <-done:
				return
			}
		}()
		return stream
	}

	take := func(done <-chan struct{}, values <-chan interface{}, num int) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			for i := 0; i < num; i++ {
				select {
				case stream <- values:
				case <-done:
					return
				}
			}
		}()
		return stream
	}

	done := make(chan struct{})
	for num := range take(done, repeatedFn(done, func() interface{} {
		return rand.Int()
	}), 10) {

		if ch, ok := num.(<-chan interface{}); ok {
			fmt.Printf("ch = %T %v \n", ch, <-ch)
		}

	}
	close(done)

}
