package main

import (
	"demo/chat31/handler"
	"net/http"
	"os"
)

//异常处理

/*
一个简单的http服务器
*/

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
			return
		}

	}
}

func main() {

	http.HandleFunc("/list/", errWrapper(handler.Handle))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
