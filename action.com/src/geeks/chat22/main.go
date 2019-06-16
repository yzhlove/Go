package main

import (
	"sync"

	"fmt"
)

//sync.Once 一个有趣的例子

func main() {

	var (
		counter int
		signale sync.Once
	)

	increment := func() {
		counter++
	}

	decrment := func() {
		counter--
	}

	signale.Do(increment)
	signale.Do(decrment)

	//counter = 1

	fmt.Println("counter = ", counter)
}
