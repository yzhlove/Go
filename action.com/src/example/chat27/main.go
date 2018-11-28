package main

import "fmt"

func main() {

	arr := []int{1,2,3}
	arr2 := []int{4,5,6,7,8}
	//temp := make([]int,0,10)

	var temp []int

	temp = arr
	fmt.Printf("%v %v %v \n",temp ,len(temp),cap(temp))

	temp = arr2
	fmt.Printf("%v %v %v \n",temp ,len(temp),cap(temp))

}
