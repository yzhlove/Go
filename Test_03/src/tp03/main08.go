package main

// 匿名函数的使用

func testFunc(f func()) {
	f()
}

func testFunc2() func(int ,int) int {
	return func(x int, y int) int {
		return ((x + y) * 2 ) << 1
	}
}


func main() {

	// 直接调用
	func (s string) {
		println(s)
	}("Hello World")

	// 赋值给变量
	add  :=  func(x,y int) int {
		return x + y
	}
	println(add(2,3))

	// 作为参数
	testFunc(func() {
		println("what are you going.")
	})

	// 作为返回值
	tf := testFunc2()
	value := tf(1,2)
	println("value = ",value)

}
