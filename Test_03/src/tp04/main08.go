package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//  修改字符串

func pp(format string,ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format,*h)
}


func main() {

	s := "Hello World"
	pp("s:%x \n",&s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte , bs : %x \n",&bs)
	pp("[]byte to string , s2 : %x \n",&s2)


	rs := []rune(s)
	s3 := string(rs)

	pp("string to []rune, rs : %x \n",&rs)
	pp("[]rnue to string ,s3 : %x \n",&s3)


}
