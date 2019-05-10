package main

import (
	"time"

	"fmt"
)

//位操作

func main() {

	ts := time.Now().UnixNano()
	fmt.Printf("%v %x %b \n", ts, ts, ts)

	var temp int64
	temp = -64
	fmt.Printf("%v %x %b \n", temp, temp, temp)

	tm := time.Now()
	fmt.Printf("%v \n", tm.YearDay())

}
