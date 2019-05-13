package main

import (
	"fmt"
	"reflect"
)

//比较map与slice

func main() {

	tempmap1 := map[int]string{1: "yzh", 2: "lcm", 3: "xjj"}
	tempmap2 := map[int]string{1: "yzh", 2: "lcm", 3: "xjj"}

	fmt.Println("map compare => ", reflect.DeepEqual(tempmap1, tempmap2))

	tempList1 := []int{1, 2, 3, 4, 5}
	tempList2 := []int{1, 2, 3, 4, 5}

	fmt.Println("slice compare => ", reflect.DeepEqual(tempList1, tempList2))

}
