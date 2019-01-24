package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//一个简单的Http请求

const method = "POST"
const url = "http://www.163.com"

func main() {

	client := &http.Client{}

	request, err := http.NewRequest(method, url, strings.NewReader("key=value"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
		return
	}

	request.Header.Add("User-Agent", "myClient")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Println(string(data))

}
