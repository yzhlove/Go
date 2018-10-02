package main

import (
	"fmt"
	"unsafe"
)

// 定义常量时,明确指定常量类型之后，赋值要确保不产生溢出

const  (
	xt,yt int  = 99,-99
	b byte = byte(xt)
	//n 	   = uint8(y)	// 八位无符号类型最多存储256，会产生溢出
	)


const (
	ptrSize = unsafe.Sizeof(uintptr(0))
	strSize = len("Hello World")
)

// 常量组中如果不指定类型和初始化值,则与上一行的非空常量右值（表达文本式）相同

func main() {
	const (
		x uint16 = 1234
		y
		s = "abc"
		z
	)

	fmt.Printf("%T ,%v\n",y,y)
	fmt.Printf("%T,%v \n",z,z)


}
