package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 切片与数组的差异

//  切片对象不允许被比较

func main() {

	var a []int
	b := []int{}

	println(a == nil,b == nil)

	fmt.Printf("a:%#v \n",(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Printf("b:%#v \n",(*reflect.SliceHeader)(unsafe.Pointer(&b)))

	fmt.Printf("a size: %d\n",unsafe.Sizeof(a))
	fmt.Printf("b size: %d\n",unsafe.Sizeof(b))


}
