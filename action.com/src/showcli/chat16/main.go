package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Line struct {
	Index int
	Count int
}

func main() {

	path := "./bytes.txt"

	lines := GetLines(path)
	fmt.Println(lines)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, v := range lines {
		if data, err := readBuf(file, int64(v.Index), v.Count); err != nil {
			fmt.Println("err ", v)
			break
		} else {
			fmt.Println(string(data))
		}
	}

}

func BackupLine(sf *os.File, lines []Line) {
	bf, err := os.Create("backupBytes.txt")
	if err != nil {
		panic(err)
	}
	defer bf.Close()
	indexMap := make(map[int]int)
	for i := 0; i < len(lines); i++ {
		indexMap[lines[i].Index] = lines[i].Count
	}
	buf := make([]byte, 10)
	for i := 0; i < len(lines); i++ {
		_, err = io.ReadFull(sf, buf)
		if err != nil {
			fmt.Println("Err:", err)
			break
		}
	}

}

func GetLines(path string) []Line {
	var (
		lines        []Line
		tempLine     Line
		index, count int
	)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR open file err.")
		return nil
	}
	defer file.Close()
	data := make([]byte, 1)
	for {
		_, err = io.ReadFull(file, data)
		if err != nil || err == io.EOF {
			break
		}
		if bytes.Compare(data, []byte("#")) == 0 {
			count = 0
			tempLine = Line{Index: index}
		}
		if bytes.Compare(data, []byte("*")) == 0 {
			tempLine.Count = count + 1
			lines = append(lines, tempLine)
		}
		count++
		index++
	}
	return lines
}

func readBuf(file *os.File, start int64, count int) ([]byte, error) {
	_, err := file.Seek(start, 0)
	if err != nil {
		return nil, err
	}
	data := make([]byte, count)
	_, err = io.ReadFull(file, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
