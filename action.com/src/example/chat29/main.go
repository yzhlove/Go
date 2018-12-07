package main

import "fmt"

func main() {

	list := make(map[int]int)
	fmt.Printf("%v \n",list)
	list[1234]++
	fmt.Printf("%v \n",list)
	list[1234]++
	fmt.Printf("%v \n",list)

}
