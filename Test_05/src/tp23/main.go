package main

import (
	"sync"
)

//错误的设计
//死锁

func main() {

	var m sync.Mutex

	//mutex不支持嵌套，会造成死锁

	m.Lock()
	{
		m.Lock()
		m.Unlock()
	}
	m.Unlock()

}
