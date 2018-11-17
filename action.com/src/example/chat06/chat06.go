package main

import (
	"fmt"
	"math/rand"
	"time"
)

//传递数组

func swap(array [1E6]int) {
	for i:= 0;i < 1E6;i++ {
		if i % 21 == 0 {
			//fmt.Println("")
		}
		//fmt.Print(array[i]," ")
	}
	fmt.Println("Done!")
}


func main() {

	var array [1E6]int

	for i:= 0;i < 1E6;i++ {
		array[i] = rand.Intn(100000) + 1
	}

	nowTime := time.Now()

	swap(array)

	fmt.Printf("%v \n",time.Now().Sub(nowTime).Seconds())


}

/*

point:0.000338791

value:0.005725469

*/