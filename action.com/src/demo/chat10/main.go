package main

import (
	"demo/chat10/examp"
	"fmt"
)

//private 与 public  技巧
func main() {

	// count := new(examp.accountNumber)	//未导出的字段，没发使用
	// *count = 100
	// fmt.Printf("count = %d \n", *count)

	//通过New来导出私有类型
	count := examp.New(100)
	fmt.Printf("count:%v \n", count)

}
