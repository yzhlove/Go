package main

import "log"

// try - catch

func main() {
	// recover捕获异常
	defer func() {
		if err := recover();err != nil {
			log.Fatalln(err)
		}
	}()
	// 抛出异常
	panic("what are you doing")
	println("exit.")	// 不会执行
}
