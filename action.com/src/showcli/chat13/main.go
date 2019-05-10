package main

import (
	"fmt"
	"io"
	"os"
)

//copy file

func main() {

	source, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer source.Close()
	distance, err := os.Create("backup.txt")
	if err != nil {
		panic(err)
	}
	defer distance.Close()
	var (
		data    []byte
		bufSize = 5
	)
	for {
		if data, err = read(source, bufSize); err != nil {
			//fmt.Printf("\n[ERROR] read %v ", err)
			break
		}
		fmt.Print(string(data))
		_, _ = distance.Write(data)
	}

	fmt.Println()
	fmt.Println("Done .")

}

func read(file *os.File, max int) ([]byte, error) {
	buf := make([]byte, max)
	_, err := io.ReadFull(file, buf)
	if err != nil {
		return nil, err
	}
	if err == io.EOF {
		return nil, fmt.Errorf("EOF")
	}
	return buf, nil
}
