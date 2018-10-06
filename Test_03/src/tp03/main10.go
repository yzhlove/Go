package main

// go 语言 闭包

func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	f := test(123)
	f()
}
