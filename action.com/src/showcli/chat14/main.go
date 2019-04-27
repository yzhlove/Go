package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 5)
	_, _ = io.ReadFull(file, buf)
	fmt.Println(string(buf))

	_, _ = io.ReadFull(file, buf)
	fmt.Println(string(buf))

}
