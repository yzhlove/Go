package main

// 多个匿名函数可能引起环境变量的改变

func testFuncChange(x int) (func(),func()) {
	return func() {
		println("Func1_x = ",x)
		x += 10
	},func() {
		println("Func2_x = ",x)
	}
}


func main() {

	a , b := testFuncChange(100)
	a()
	b()

}