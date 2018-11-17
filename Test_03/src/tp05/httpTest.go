//HTTP 测试 (GET)
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("http://www.baidu.com")
	if nil != err {
		log.Fatalln("ERR:", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		log.Fatalln("Read File ERR:", err)
	}
	fmt.Println("http response:", string(body))
}
