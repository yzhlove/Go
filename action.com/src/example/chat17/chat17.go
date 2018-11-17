package main

import (
	"fmt"
	"math/rand"
)

//map删除操作

func main() {

	hashMap := make(map[int]string)

	colors := [5]string{"Red","Orange","Yellow","Green","Blue"}

	for index,value := range colors {
		hashMap[index + 1] = value
	}

	fmt.Printf("%+v \n",hashMap)

	//删除map中的元素
	delete(hashMap,rand.Intn(5) + 1)

	fmt.Printf("%+v \n",hashMap)

}