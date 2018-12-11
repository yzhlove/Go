package main

import (
	"fmt"
	"math/rand"
)

func splitScore(amount, count int) []int {

	sum := 0
	list := make([]int, count, count)
	avg := int(amount / count)
	fmt.Printf("avg: %v \n", avg)

	if avg < 2 {
		avg = 1
	}

	for i := 0; i < count; i++ {
		temp := 0
		if sum <= amount {
			temp = rand.Intn(avg) + rand.Intn(2) + int(avg/2)
			sum += temp
		}
		list[i] = temp
	}
	for index := range list {
		if sum > amount {
			break
		}
		list[index] += avg
		sum += avg
	}

	sumCount := 0
	for i := 0; i < count; i++ {
		sumCount += list[i]
	}
	fmt.Printf("sumCount = %v \n", sumCount)

	return list
}

func main() {

	fmt.Printf("%v \n", splitScore(10, 15))
	fmt.Printf("%v \n", splitScore(20, 15))
	fmt.Printf("%v \n", splitScore(15, 15))
	fmt.Printf("%v \n", splitScore(50, 15))
	fmt.Printf("%v \n", splitScore(100, 15))

}
