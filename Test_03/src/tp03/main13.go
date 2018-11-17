package main


// 闭包个性

// 返回一个函数数组
func test() []func() {

	var s []func()

	for i := 0;i < 5;i++ {
		x := i
		s = append(s, func() {
			println(&x,x)
		})
	}
	return s
}


func main() {

	for _,f := range test() {
		f()
	}

}
