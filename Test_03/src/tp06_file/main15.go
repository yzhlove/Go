package main

import (
	"time"
)

//空结构可作为通道类型，用于事件通知
func main() {

	exit := make(chan struct{})

	go func() {
		time.Sleep(100)
		println("go func")
		exit <- struct{}{}
	}()

	<-exit
	println(".end")
}
