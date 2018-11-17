package main

import (
	"fmt"
	"math"
)

//Go 取整数

func main() {

	var a = 23.456
	var b = 23.789

	fmt.Println(math.Ceil(23.456))
	fmt.Println(math.Ceil(23.789))
	fmt.Println(math.Floor(23.456))
	fmt.Println(math.Floor(23.789))
	fmt.Println(int(a))
	fmt.Println(int(b))
}
