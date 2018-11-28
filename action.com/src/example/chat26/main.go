package main

import "fmt"

//找出最长子串

func maxLength(s string) int {
	lastString := make(map[rune]int)
	maxLength , start := 0,0
	for index,value := range []rune(s) {
		lastI , ok := lastString[value]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if index - start + 1 > maxLength {
			maxLength = index - start + 1
		}
		lastString[value] = index
	}
	return maxLength
}

func main() {
	fmt.Println("length = " , maxLength("abcdefghij"))
}
