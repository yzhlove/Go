package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//一个简单的CURL

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example <url>")
		os.Exit(-1)
	}
}

func main() {

	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//复制网络流输出到屏幕
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err.Error())
	}
}
