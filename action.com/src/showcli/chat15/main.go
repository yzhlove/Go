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

	fmt.Println("-------------------------------")

	tf, err := os.Create("tempBytes.txt")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	//移动文件指针
	_, _ = file.Seek(0, 0)
	var (
		data []byte
		n    int
	)
	var countIndex int
	for i := 0; i < len(lines); i++ {
		tmp := lines[i]
		data, n, err = readIndex(file, tmp.Index+tmp.Length-countIndex+1)
		_, _ = tf.Write(data)
		countIndex += n
		fmt.Printf("countIndex = %v line:%v read:\n[%v] \n", countIndex, tmp, string(data))
	}
	data, n, err = readIndex(file, 128)
	fmt.Println("last n = ", n)
	_, _ = tf.Write(data)
	fmt.Println("Done .")
}

type Line struct {
	Index  int
	Length int
}

func readIndex(file *os.File, byteLength int) ([]byte, int, error) {
	buf := make([]byte, byteLength)
	var n int
	var err error
	if n, err = io.ReadFull(file, buf); err != nil {
		return nil, 0, errors.New("EOF")
	}
	return buf, n, nil
}

func readByte(file *os.File) ([]byte, error) {
	buf := make([]byte, 1)
	if _, err := io.ReadFull(file, buf); err != nil || err == io.EOF {
		return nil, errors.New("EOF")
	}
	return buf, nil
}
