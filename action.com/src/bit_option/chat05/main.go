package main

import (
	"fmt"
	"time"
)

func main() {

	a := time.Now().Unix()
	time.Sleep(2 * time.Second)
	b := time.Now().Unix()

	fmt.Println(a - b)

}
