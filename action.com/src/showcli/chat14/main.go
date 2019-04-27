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
	buf := make([]byte, 1024)
	n, _ := io.ReadFull(file, buf)
	fmt.Println("n = ", n)
	fmt.Println(string(buf))

	//_, _ = io.ReadFull(file, buf)
	//fmt.Println(string(buf))

}
