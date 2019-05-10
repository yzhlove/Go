package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	path := "./test.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("FileName = ", file.Name())

	var index int64
	var length int
	if index, err = file.Seek(0, 1); err != nil {
		panic(err)
	}

	data := make([]byte, 5)
	if length, err = io.ReadFull(file, data); err != nil || err == io.EOF {
		panic(err)
	}
	fmt.Printf("length = %v index = %v \n", length, index)
	fmt.Println(string(data))

	//_, _ = file.Seek(5, 1)
	_, _ = io.ReadFull(file, data)
	fmt.Println(string(data))

	_, _ = file.Seek(3, 0)
	_, _ = io.ReadFull(file, data)
	fmt.Println(string(data))

}
