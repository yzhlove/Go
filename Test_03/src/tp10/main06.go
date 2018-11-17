package main

import (
	"sync"
)

//匿名成员

//Buff 缓冲
type Buff struct {
	sync.Mutex
	buf [1024]byte
}

func main() {

	bf := Buff{}

	//调用匿名成员
	bf.Lock()
	defer bf.Unlock()

}
