package main

import "fmt"

//切片测试

func main() {

	list := []int{1,3,5,7,9,11,13,15,17,19}

	var index int

	for _, tmp := range list {
		fmt.Printf("value = %v \n",tmp)
		if tmp == 13 {
			index = tmp
			break
		}
	}

	fmt.Printf("index = %v \n",index)

}
