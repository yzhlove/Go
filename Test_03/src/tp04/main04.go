package main

// 跨行

func main() {

	s := "ab" +		// 跨行时，操作符必须在上一行结尾
		"cd"
	println(s == "abcd")
	println(s > "abc")

}

