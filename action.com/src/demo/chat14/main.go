package main

import (
	"fmt"
	"runtime"
)

func main() {

	cpu := runtime.NumCPU()
	fmt.Printf("cpu:%v \n", cpu)

}
