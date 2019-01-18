package main

import (
	"fmt"
	"runtime"
)

//宕机恢复

type paicContext struct {
	function string
}

//以保护的形式运行一个函数
func proteRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("begin ...")
	proteRun(func() {
		panic(&paicContext{
			function: "手动触发panic",
		})
	})

	proteRun(func() {
		fmt.Println("宕机前")
		var a *int
		//a = new(int)	// NULL指针复值报错
		*a = 1
		fmt.Println("宕机后")
	})
	fmt.Println("end ...")
}
