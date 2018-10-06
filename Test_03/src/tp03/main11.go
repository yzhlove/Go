package main


// 闭包解析

 func test(x int) func() {

 	println("&x = ",&x,x)

 	return func () {
 		println("& xx = ",&x,x)
	}

 }


func main() {

	f := test(123)
	f()

}