package main

import "log"

//import "log"

// try - catch 执行流程

func test() {
	defer println("what are you doing")
	defer println("are you kidding me")

	// 无论是否调用recover，所有的defer都会被执行
	panic("i love you")
}

func main() {

	defer func() {
		log.Fatalln(recover())
	}()

	test()

}

/*
are you kidding me
what are you doing
2018/10/06 11:05:30 i love you
*/