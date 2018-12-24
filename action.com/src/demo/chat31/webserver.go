package main

import (
	"demo/chat31/handler"
	"net/http"
)

//异常处理

/*
一个简单的http服务器
*/

func main() {

	http.HandleFunc("/list/", handler.Handle)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
