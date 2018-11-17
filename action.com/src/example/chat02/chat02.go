package main

import "fmt"

//指针数组

func main() {

	array := [5]*int{0:new(int),1:new(int)}
	*array[0] = 100
	*array[1] = 200

	for _,v := range array {
		if v == nil {
			continue
		}
		fmt.Printf("%v \n",*v)
	}

}


