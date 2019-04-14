package main

import (
	"fmt"
	"io"
	"os"
)

//文件备份

func main() {

	var (
		sourceFile, newFile *os.File
		err                 error
		sourcePath          = "/Users/love/redis_filter.txt"
		bytes               int64
	)
	if sourceFile, err = os.Open(sourcePath); err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	if newFile, err = os.Create(sourcePath + "_copy"); err != nil {
		panic(err)
	}

	defer newFile.Close()

	if bytes, err = io.Copy(newFile, sourceFile); err != nil {
		panic(err)
	}

	fmt.Println("copy bytes = ", bytes)

	//将文件刷写到硬盘
	if err = newFile.Sync(); err != nil {
		panic(err)
	}

	fmt.Println("文件刷写成功")

}
