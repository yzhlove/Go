package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

//Handle 处理函数
func Handle(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) < 0 {
		return userError("must " + prefix + " is head")
	}
	path := request.URL.Path[len(prefix):]
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
