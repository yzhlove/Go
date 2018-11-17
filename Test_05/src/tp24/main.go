package main

import (
	"fmt"
	"sync"
)

//一个错误的锁设计

type cache struct {
	sync.Mutex
	data []int
}

func (c *cache) count() int {
	// c.Lock()
	n := len(c.data)
	// c.Unlock()
	return n
}

func (c *cache) get() int {
	c.Lock()
	defer c.Unlock()

	var d int
	if n := c.count(); n > 0 { //重复锁，导致死锁，错误
		d = c.data[0]
		c.data = c.data[1:]
	}
	return d
}

func main() {
	c := cache{
		data: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(c.get())
}
