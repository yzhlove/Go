package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//切片实验

func main() {

	slice := make([]string,2,5)

	fmt.Printf("%#v \n",slice)

	slice[0] = "hello"
	slice[1] = "world"

	fmt.Printf("%#v \n",slice)

	for i := 0;i < len(slice);i++ {
		slice[i] = strconv.FormatInt(int64(i),10)
	}

	fmt.Printf("%#v \n",slice)

	//panic: runtime error: index out of range
	/*for i:= 0;i < cap(slice);i++ {
		slice[i] = strconv.FormatInt(int64(i * 10 + 1),10)
	}

	fmt.Printf("%#v \n",slice)*/

	capLength := cap(slice)
	for i:= 0;i < capLength;i++ {
		slice = append(slice,strconv.FormatInt(int64(i * 10 + 1),10))
	}

	fmt.Printf("%#v \n",slice)

	slice2 := slice[2:6]

	fmt.Printf("%#v \n",slice2)

	//range copy slice无法修改
	for index,value := range slice2 {
		fmt.Printf("[%v:%v] \n",index,value)
		value = strconv.FormatInt(int64(rand.Intn(index*10 + 1) + 1),10)
	}

	fmt.Printf("slice:%#v \nslice2:%#v \n",slice,slice2)

	fmt.Println("----------------------------------")

	//修改slice
	for index , value := range slice2 {
		fmt.Printf("[%v:%v] \n",index,value)
		slice2[index] = strconv.FormatInt(int64(rand.Intn(index*10 + 1) + 1),10)
	}

	fmt.Printf("slice:%#v \nslice2:%#v \n",slice,slice2)

}