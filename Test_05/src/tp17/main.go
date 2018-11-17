package main

import (
	"fmt"
	"time"
)

//通道实现信号量

type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p:
	default:
		v = make([]byte, 10)
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b:
	default:
	}
}

func main() {

	pl := newPool(3)
	pl.put([]byte{1, 2, 3, 4, 5})

	for v := range pl.get() {
		fmt.Printf("%v \n", v)
	}

	time.Sleep(time.Second)

}
