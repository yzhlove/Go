package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

//Handle 处理函数
func Handle(writer http.ResponseWriter, request *http.Request) error {

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	//读取文件
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if _, err = writer.Write(all); err != nil {
		return err
	}
	return nil
}
