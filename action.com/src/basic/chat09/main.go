package main

import "fmt"

//copy 与 reference

func main() {

	const number = 1000
	srcData := make([]int, number)
	for i := 0; i < number; i++ {
		srcData[i] = i
	}
	//引用数据
	refData := srcData
	copyData := make([]int, number)
	copy(copyData, srcData)

	//修改原始数据
	srcData[0] = 999
	fmt.Printf("srcData[0] = %d \n", srcData[0])
	fmt.Printf("refData[0] = %d \n", refData[0])

	fmt.Printf("copyData: %d %d \n", copyData[0], copyData[number-1])

	copy(copyData, srcData[4:6])
	fmt.Printf("len = %d \n", len(copyData))
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", copyData[i])
	}

}
