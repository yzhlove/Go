package main

import (
	"fmt"
)

//map

func main() {

	hashMap := make(map[string][]int)
	hashMap["A"] = []int{10, 20}
	hashMap["B"] = []int{20, 30}
	hashMap["C"] = []int{30, 40}
	hashMap["D"] = []int{40, 50}
	hashMap["E"] = []int{50, 70}

	for key, array := range hashMap {
		for _, arrayValue := range array {
			fmt.Println("key:", key, " value:", arrayValue)
		}
		fmt.Println()
	}

}
