package main

import (
	"fmt"
	"time"
)

const (
	HighClear = 0x0000FFFF
	LowClear  = 0xFFFF0000
)

func main() {

	ct := time.Now()

	var temp uint32
	fmt.Printf("%v -> %b \n", temp, temp)

	//set low
	ts := uint32(ct.YearDay())
	temp |= ts
	fmt.Printf("%v -> %b \n", temp, temp)

	//set high
	var count uint32 = 321
	temp |= count << 16
	fmt.Printf("%v -> %b \n", temp, temp)

	//get low
	fmt.Printf("%v -> %b \n", temp&0xFFFF, temp&0xFFFF)

	//get high
	fmt.Printf("%v -> %b \n", temp>>16, temp>>16)
	fmt.Println("-------------------")
	//高位清零
	temp &= HighClear
	fmt.Printf("%v -> %b \n", temp&0xFFFF, temp&0xFFFF)
	fmt.Printf("%v -> %b \n", temp>>16, temp>>16)

	fmt.Println("-------------------")

	count = 2456
	temp |= count << 16
	fmt.Printf("%v -> %b \n", temp>>16, temp>>16)

}
