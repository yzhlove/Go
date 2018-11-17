package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTP POST 测试

func main() {

	request := make(map[string]string)
	request["sid"] = "123456sabcdefg"
	request["url"] = "www.baidu.com"

	requestByte, err := json.Marshal(request)
	if nil != err {
		log.Fatalln("JsonERR:", err)
		return
	}

	fmt.Println("jsonData:", string(requestByte))

	requestURL := "http://jump.com/test/test/index"
	responseData, err := http.Post(requestURL, "application/json;charset=utf-8", bytes.NewBuffer(requestByte))
	defer responseData.Body.Close()
	if nil != err {
		fmt.Println(err.Error())
	}
	content, err := ioutil.ReadAll(responseData.Body)
	if nil != err {
		fmt.Println(err.Error())
	}

	fmt.Println(string(content))

}
