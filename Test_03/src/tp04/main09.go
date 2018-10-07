package main

import (
	"fmt"
	"unsafe"
)

// 修改字符串

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {

	bs := []byte("Hello World")
	s := toString(bs)

	fmt.Printf("bs: %x \n",bs)
	fmt.Printf("s: %x \n",&s)



}