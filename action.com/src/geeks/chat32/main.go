package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var stdout bytes.Buffer
	defer stdout.WriteTo(os.Stdout)
	insStream := make(chan int, 4)
	go func() {
		defer close(insStream)
		defer fmt.Fprintf(&stdout, "Done -> \n")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdout, "Send:%d \n", i)
			insStream <- i
		}
	}()

	for integer := range insStream {
		fmt.Fprintf(&stdout, "Receive:%d \n", integer)
	}
}
