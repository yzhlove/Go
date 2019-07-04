package main

import "fmt"

func main() {

	var a uint32 = 10
	var b uint32
	var c uint32 = 20

	b |= a << 16 //high
	b |= c
	fmt.Println("high b = ", b>>16)
	fmt.Println("low b = ", b&0xFFFF)

	var e uint32 = 23456
	var f uint32 = 1314

	b >>= 32

	b |= e << 16
	b |= f
	fmt.Println("high b = ", b>>16)
	fmt.Println("low b = ", b&0xFFFF)

}
