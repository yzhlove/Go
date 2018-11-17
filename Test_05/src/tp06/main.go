package main

import (
	"fmt"
	"math/rand"
)

// 通道用完之后要及时close掉，否则会造成死锁

func main() {

	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for x := range c {
			fmt.Println("x = ", x)
		}
	}()

	for i := 1; i <= 10; i++ {
		c <- rand.Intn(100) + 50
	}

	close(c)
	<-done

}
