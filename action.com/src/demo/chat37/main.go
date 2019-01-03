package main

import "fmt"

func main() {

	str := "我爱我的祖国"

	for k, v := range []rune(str) {
		fmt.Printf("%v %c %c\n", k, v, str[k])
	}

}
