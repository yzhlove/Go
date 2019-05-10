package main

import (
	"time"

	"fmt"
)

func main() {

	var temp uint32

	ts := time.Now()
	fmt.Printf("%b \n", temp)
	temp |= uint32(ts.YearDay())
	fmt.Printf("%v %b \n", temp, temp)
	fmt.Printf("%v %b \n", 0xFFFF, 0xFFFF)

	count := 256
	temp = uint32(count) &^ temp
	fmt.Printf("%v %b \n", temp, temp)

}
