package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

//请求手机网站

const (
	url         = "http://www.imooc.com"
	requestHead = "User-Agent"
	userAgent   = `Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Mobile Safari/537.36`
)

func main() {
	var (
		request  *http.Request
		err      error
		response *http.Response
		buffer   []byte
	)
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		panic(err)
	}
	request.Header.Add(requestHead, userAgent)
	//使用默认的 client
	//if response, err = http.DefaultClient.Do(request); err != nil {
	//	panic(err)
	//}

	//自定义client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect URL:", req)
			// 返回nil允许重定向  返回err不允许重定向
			return nil
		},
	}

	if response, err = client.Do(request); err != nil {
		panic(err)
	}

	if buffer, err = httputil.DumpResponse(response, true); err != nil {
		panic(err)
	}
	fmt.Println(string(buffer))
}
