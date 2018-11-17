package main

import "fmt"

// 字符串

func main() {

	s := "雨痕\x61\142\u0041"
	fmt.Printf("%s\n",s)
	fmt.Printf("% x,len:%d \n",s , len(s))

}

/*
雨痕abA
e9 9b a8 e7 97 95 61 62 41,len:9
*/