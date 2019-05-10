package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	stopper := time.NewTimer(time.Second * 2)

	var i int

EXIT:
	for {
		select {
		case <-stopper.C:
			fmt.Println("stop")
			break EXIT
		case <-ticker.C:
			i++
			fmt.Println("tick count:", i)
		}
	}
	fmt.Println("done.")

}
