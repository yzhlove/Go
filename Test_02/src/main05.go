package main

// defer 暂时理解为C++析构函数

func Test(a,b int)  {
	defer println("what are you doing")

	println(a / b)
}

func returnFunction() func() {
	return func(){
		println("This is func")
	}
}

func Test2(c func())  {
	defer c()
	println("12345")
}

func main() {

	a ,b := 100,2
	Test(a,b)

	println("--------------------")

	f := returnFunction()
	Test2(f)

}