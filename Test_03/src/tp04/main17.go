package main

import "fmt"

// 指针数组

func main() {

	x , y := 10,20

	a := [...]*int{&x,&y}	// 指针数组

	p := &a	// 数组指针

	fmt.Printf("%T,%v\n",a,a)
	fmt.Printf("%T,%v\n",p,p)

}
