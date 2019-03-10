package main

import "sync/atomic"

var value int32

func SetValue(detal int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+detal) {
			break
		}
	}
}

func main() {
	SetValue(123)
}
