package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Handle(writer http.ResponseWriter, request *http.Request) {

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		http.Error(writer,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	//读取文件
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if _, err = writer.Write(all); err != nil {
		panic(err)
	}
}
