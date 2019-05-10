package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//文件拷贝

func main() {

	path := "/Users/love/redis_filter.txt"

	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Ext(path))

	fileInfo, err := os.Stat(filepath.Dir(path))
	if err != nil {
		panic(err)
	}
	if fileInfo.IsDir() {
		fmt.Printf("是个文件夹")
	}
}
