package main

// 闭包

/*
test => func()
*/
func test(x int) func() {
	return func() {
		println(x)
	}
}

func main()  {
	x := 100
	f := test(x)
	f()
}