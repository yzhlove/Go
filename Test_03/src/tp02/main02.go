package main

//  指针操作

func main() {

	x := 100;
	var p *int = &x
	*p += 20

	println(p,*p)


}
