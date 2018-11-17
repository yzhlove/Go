package main

import "fmt"

// 字符串遍历

func main() {

	s := "雨痕"

	for i:= 0; i < len(s);i++ {
		fmt.Printf("%d:[%c]\n",i,s[i])
	}

	for i , c := range s {
		fmt.Printf("%d:[%c]\n",i,c)
	}

}
