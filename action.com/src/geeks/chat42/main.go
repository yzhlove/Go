package main

import (
	"fmt"
	"net/http"
	"time"
)

//错误处理

type Result struct {
	Err  error
	Resp *http.Response
}

func main() {

	urls := []string{"http://www.baidu.com", "a", "b", "c"}
	done := make(chan struct{})
	responseQueue := respQueue(done, urls...)

	for resp := range responseQueue {
		if resp.Err != nil {
			fmt.Println("Resp Err:", resp.Err)
			continue
		}
		fmt.Println("Resp Status:", resp.Resp.Status)
	}
	close(done)
	time.Sleep(time.Second)
	fmt.Println("Done .")
}

func respQueue(done <-chan struct{}, urls ...string) chan Result {
	responseQueue := make(chan Result)
	go func() {
		defer fmt.Println("exit .")
		defer close(responseQueue)
		for _, url := range urls {
			resp, err := http.Get(url)
			select {
			case responseQueue <- Result{Err: err, Resp: resp}:
			case <-done:
				return
			}
		}
	}()
	return responseQueue
}
