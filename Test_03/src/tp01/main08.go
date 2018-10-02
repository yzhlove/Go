package main

import "fmt"

//  go枚举实现

const (
	xx = iota 	// 0
	yx
	zx
)


const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
)


func main() {

	fmt.Printf("%T,%v \n",xx,xx)
	fmt.Printf("%T,%v \n",yx,yx)
	fmt.Printf("%T,%v \n",zx,zx)
	//fmt.Printf("%T,%v \n",_)
	fmt.Printf("%T,%v \n",KB,KB)
	fmt.Printf("%T,%v \n",MB,MB)
	fmt.Printf("%T,%v \n",GB,GB)


}
