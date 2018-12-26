package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

//一个简单的CURL

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//同时向 (标准输出和文件写数据)
	dest := io.MultiWriter(os.Stdout, file)
	//读取响应的内容，并写入（标准输出和文件）
	io.Copy(dest, resp.Body)
	if err := resp.Body.Close(); err != nil {
		log.Println(err)
	}
}
