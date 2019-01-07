package main

import (
	"fmt"
)

//函数接口实现

type Invoker interface {
	Call(interface{})
}

type Model struct{}

func (m *Model) Call(p interface{}) {
	fmt.Println("format struct ", p)
}

type FuncCaller func(interface{})

func (fn FuncCaller) Call(p interface{}) {
	fn(p)
}

func main() {
	var invoker Invoker

	model := new(Model)
	invoker = model
	invoker.Call("hello")

	//将匿名函数强制转换为FuncCaller类型
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function ", v)
	})
	invoker.Call("Hello World")
}
