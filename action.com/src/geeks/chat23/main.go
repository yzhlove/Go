package main

import "sync"

// 一个错误的例子
// 死锁

func main() {

	var (
		onceA, onceB sync.Once
		initA, initB func()
	)

	initA = func() {
		onceB.Do(initB)
	}

	initB = func() {
		onceB.Do(initA)
	}

	onceA.Do(initA)

}
