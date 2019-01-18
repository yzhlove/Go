package main

import (
	"fmt"
	"os"
)

//获取文件的详细信息

func fileSize(filename string) int64 {

	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0
	}
	size := info.Size()
	name := info.Name()
	ts := info.ModTime()
	mode := info.Mode()

	fmt.Printf("size:%v name:%v ts:%v mode:%v \n", size, name, ts, mode)
	return size
}

func main() {

	fileSize("./access.log")

}
