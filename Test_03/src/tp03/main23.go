package main

import "runtime/debug"

// 利用 try - catch 做一些处理


func test(x,y int) {

	z := 0

	func() {
		defer func() {
			if recover() != nil {
				debug.PrintStack()
				z = 0
			}
		}()

		z = x / y

	}()

	println("x / y = ",z)

}

func main() {
	test(5,0)
}