package main

import "fmt"

func main() {

	//map的使用
	keyMap := make(map[int][]string)
	keyMap[1] = []string{"a", "b", "c"}

	fmt.Printf("%T %v \n", keyMap[1], keyMap[1])
	fmt.Printf("%T %v \n", keyMap[2], keyMap[2])

}
