package main

// 延迟调用

func main() {

	x ,y := 1,2
	defer func(a int) {			// 仅注册，在main函数结束前调用
		println("defer x,y = ",a,y)
	}(x)	// x 的值会被缓存起来，x = 1

	x += 100
	y += 100
	println("x = " , x," y = ",y)

}
