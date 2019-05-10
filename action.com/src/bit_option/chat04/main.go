package main

import (
	"fmt"
	"math"
)

func main() {

	var a float32 = 4.5
	var b float32 = 5.86

	fmt.Println(a * b)

	fmt.Printf("%v %v %v \n", math.Trunc(float64(a*b)), math.Ceil(float64(a*b)), uint32(a*b))

}
