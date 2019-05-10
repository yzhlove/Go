package main

import "time"

func main() {

	flag := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		close(flag)
	}()

	<-flag

}
