package main

import (
	"fmt"
	"time"
)

func main() {

	var index int
LOOP:
	for {
		select {
		case <-time.NewTimer(1 * time.Second).C:
			fmt.Println("Hello world")
			index++
			if index == 3 {
				break LOOP
			}
		}
	}

	fmt.Println("Done .")
}
