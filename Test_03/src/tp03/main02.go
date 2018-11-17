package main


// 关于传递指针

func test(p **int) {
	x := 100
	*p = &x
}

func main() {
	var p *int
	test(&p)
	println("*p = ",*p)
}
