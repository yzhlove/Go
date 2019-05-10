package main

import (
	"demo/chat31/handler"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

//异常处理

/*
一个简单的http服务器
*/

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Exception:%v \n", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			if usrErr, ok := err.(userErr); ok {
				http.Error(writer, usrErr.Message(), http.StatusBadRequest)
				return
			}
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

type userErr interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/list/", errWrapper(handler.Handle))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
