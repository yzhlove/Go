package main

import (
	"fmt"
)

func main()  {

	x := make([]int , 0 , 5)	// 创建容量为5的切片

	for i := 0; i < 10;i++ {
		x = append(x,i)
	}

	fmt.Println(x)

	fmt.Println("-----------------")

	mmp := make(map[string]int)  // 创建字典对象
	mmp["a"] = 1

	y ,ok := mmp["a"]
	fmt.Println(y,ok)

	delete(mmp,"a")





}


