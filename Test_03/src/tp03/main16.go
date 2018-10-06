package main

// 延迟调用

func testDefer() (z int) {
	defer func() {
		println("defer:",z)
		z += 100
	}()
	return 100
}

func main() {

	println("test:",testDefer())

}

/*
defer: 100
test: 200
*/