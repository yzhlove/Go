package main

import (
	"fmt"
	"os"
	"time"
)

//timeout与timetick

func main() {

	go func() {
		select {
		case <-time.After(time.Second * 5):
			fmt.Println("timeOut...")
			os.Exit(0)
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil)

}
