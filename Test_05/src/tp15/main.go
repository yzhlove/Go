package main

import (
	"fmt"
)

//select 利用default做默认处理

func main() {

	done := make(chan struct{})
	data := []chan int{
		make(chan int, 3),
	}

	go func() {
		defer close(done)
		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i:
			default:
				data = append(data, make(chan int, 3))
			}
		}
	}()

	<-done

	for i := 0; i < len(data); i++ {
		c := data[i]
		close(c)
		for x := range c {
			fmt.Println("x = ", x)
		}
	}

}
