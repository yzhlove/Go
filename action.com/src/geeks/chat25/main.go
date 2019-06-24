package main

import (
	"fmt"
	"sync"
)

//pool 的使用

type instance struct {
	Name string
	Age  int
}

func main() {

	pool := sync.Pool{
		New: func() interface{} {
			return instance{}
		},
	}

	ins := pool.Get()
	ins = instanceInit("yzh", 18)

	fmt.Printf("%+v \n", ins)

	pool.Put(ins)

	ins2 := pool.Get()
	ins2 = instanceInit("xjj", 18)

	fmt.Printf("%+v \n", ins2)

	pool.Put(ins2)

	fmt.Printf("%+v \n", pool.Get())

	pool.Put(ins)

}

func instanceInit(name string, age int) instance {
	return instance{Name: name, Age: age}
}
