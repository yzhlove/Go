package main

import (
	"fmt"
	"time"
)

//协程初步

func main() {

	var a [10]int

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//a[i]++
				//runtime.Gosched()
				a[i]++
				fmt.Printf("ID -> %d : %d \n", i, a[i])
			}
		}(i)
	}
	time.Sleep(1 * time.Minute)
	fmt.Println(a)
}
