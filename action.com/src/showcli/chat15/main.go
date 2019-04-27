package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

//bytes.txt

func main() {

	path := "./bytes.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	index := 0
	count := 0
	var lines []Line
	var tempLine Line
	for {
		data, err := readByte(file)
		if err != nil {
			break
		}
		fmt.Printf("%v :%v \n", index, string(data))
		if bytes.Compare(data, []byte("#")) == 0 {
			count = 0
			tempLine = Line{Index: index}
		}
		if bytes.Compare(data, []byte("*")) == 0 {
			tempLine.Length = count
			lines = append(lines, tempLine)
		}
		count++
		index++
	}
	fmt.Println()
	for _, v := range lines {
		fmt.Println(v)
	}

	fmt.Println("Done .")

}

type Line struct {
	Index  int
	Length int
}

func readByte(file *os.File) ([]byte, error) {
	buf := make([]byte, 1)
	if _, err := io.ReadFull(file, buf); err != nil || err == io.EOF {
		return nil, errors.New("EOF")
	}
	return buf, nil
}
