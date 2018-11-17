package main

import "fmt"

//go语言并不认同type A int 与 int是同一种类型
//go语言不做隐式转换

type MAXTYPE int64

func main() {

	var a MAXTYPE = 1000

	 //a = int64(1000)	//因为go不支持隐式准换，所以编译器会提示:cannot use int64(1000) (type int64) as type MAXTYPE in assignment

	fmt.Printf("%T , %v \n",a,a)

}

