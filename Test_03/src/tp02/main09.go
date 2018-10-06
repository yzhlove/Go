package main

import "fmt"

// range 使用

func main() {

	data := [3]int{10,20,30}

	for i ,x := range data {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x:%d data:%d \n",x,data[i])
	}

	data2 := [3]int {10,20,30}
	for i,x := range data2[:] {
		if i == 0 {
			data2[0] += 100
			data2[1] += 200
			data2[2] += 300
		}
		fmt.Printf("x:%d data:%d \n",x,data2[i])
	}


}
