package main

// 多变量赋值

func main() {

	x,y := 1,2
	x,y = y + x, x + 2	// 先计算 y+x 然后计算 x+2 最后为x,y赋值

	println("x = ",x," y = ",y)


}
