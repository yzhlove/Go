package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	var (
		wg   sync.WaitGroup
		urls = []string{
			"https://www.sina.com.cn",
			"https://www.baidu.com",
			"https://www.bing.com",
		}
	)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println("url:", url, " err:", err)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}
