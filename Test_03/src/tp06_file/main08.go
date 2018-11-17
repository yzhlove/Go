package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	time.Sleep(3)
	overTime := time.Now().Sub(currentTime)
	fmt.Printf("%T,%+v",overTime,overTime)

}
