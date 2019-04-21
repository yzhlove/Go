package main

import (
	"fmt"
	"time"
)

func main() {

	inChan := make(chan int, 2)

	go func() {
		for {
			select {
			case v, ok := <-inChan:
				if ok {
					fmt.Println("value = ", v)
				} else {
					fmt.Println("close .")
					goto LOOP
				}
			}
		}
	LOOP:
		fmt.Println("Thread Done .")
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 20; i++ {
		inChan <- i * 2
		inChan <- i * 4
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Send Done .")
	time.Sleep(200 * time.Millisecond)
	close(inChan)
	fmt.Println("Close Done .")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Done .")

}
