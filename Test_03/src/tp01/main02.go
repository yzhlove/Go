package main

// 查看变量的地址

var x int = 100

func main() {

	println("x addredd = ",&x, " x value = ",x);

	//x := "abc"
	//println("x addredd = ",&x, " x value = ",x);

	{
		x := "abc"
		println("x addredd = ",&x, " x value = ",x);
	}

	println("x addredd = ",&x, " x value = ",x);

}
