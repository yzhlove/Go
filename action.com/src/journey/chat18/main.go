package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

//Http服务器

//一个简答的Http客户端

func getRequest(url string) ([]byte, error) {
	if response, err := http.Get(url); err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if result, err := httputil.DumpResponse(response, true); err != nil {
			return nil, err
		} else {
			return result, nil
		}
	}
}

func main() {
	var (
		buffer []byte
		err    error
	)

	if buffer, err = getRequest("http://www.imooc.com"); err != nil {
		fmt.Printf("Request Error!")
	} else {
		fmt.Println(string(buffer))
	}
}
