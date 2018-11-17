package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Go生成随机数

func main() {

	rand.Seed(time.Now().Unix())
	for i := 0; i < 50; i++ {
		value := rand.Intn(10)
		fmt.Println(value)
	}

}
