package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	s := "abcdefg"

	s1 := s[:3]		// 从头开始，到3结束
	s2 := s[1:4]	// 从1开始,到4结束
	s3 := s[2:]		// 从2开始，到后面全部内容

	println(s1,s2,s3)

	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s1)))

}


