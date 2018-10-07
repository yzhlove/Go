package main

import "fmt"

// 使用缩影访问数组

func main() {

	s := "abc"
	println(s[1])
	fmt.Printf("%T,%x\n",s,s)
	//println(&s[1])			// 不允许这样取地址

	for i:= 0;i < len(s);i++ {
		println(s[i])
	}

}
