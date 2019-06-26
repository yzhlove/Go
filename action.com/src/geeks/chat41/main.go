package main

import (
	"fmt"
	"net/http"
)

//错误处理

func main() {

	urls := []string{"http://www.baidu.com", "http://www.aabbcc.com"}
	done := make(chan struct{})
	for resp := range respQueue(done, urls...) {
		fmt.Println("resp =>", resp.Status)
	}
	close(done)
	fmt.Println("Done .")
}

func respQueue(done <-chan struct{}, urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	go func() {
		defer fmt.Println("exit .")
		defer close(responses)
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("URL:%v Err:%v \n", url, err)
				continue
			}
			select {
			case responses <- resp:
			case <-done:
				return
			}
		}
	}()
	return responses
}
