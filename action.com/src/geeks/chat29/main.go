package main

import (
	"fmt"
	"time"
)

//从已关闭的channel读取数据

func main() {

	stringStream := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			value, ok := <-stringStream
			fmt.Printf("value:%v status:%v \n", value, ok)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		stringStream <- "hello world"
	}

	close(stringStream)
	time.Sleep(time.Second * 15)
	fmt.Println("Done .")
}
