package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

//模拟Nginx access.log

func main() {

	file, err := os.OpenFile("./access.log", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Open File Err:%s", err.Error()))
	}
	defer file.Close()

	for {
		for i := 0; i < 4; i++ {
			now := time.Now()
			paths := []string{"/foo", "/bar", "/baz", "/qux", "/foo", "/bar", "/bar", "/bar"}
			path := paths[rand.Intn(len(paths))]
			requestTime := rand.Float64()
			if path == "/foo" {
				requestTime = requestTime + 1.4
			}

			scheme := "http"
			if now.UnixNano()/1000%2 == 1 {
				scheme = "https"
			}
			dateTime := now.Format("02/Jan/2006:15:04:05")
			code := 200
			if now.Unix()%10 == 1 {
				code = 500
			}
			bytesSend := rand.Intn(1000) + 500
			if path == "/foo" {
				bytesSend = bytesSend + 1000
			}
			line := fmt.Sprintf("172.0.0.12 - - [%s +0000] %s \"GET %s HTTP/1.0\" %d %d \"-\" \"KeepAliveClient\" \"-\" - %.3f\n", dateTime, scheme, path, code, bytesSend, requestTime)
			_, err := file.Write([]byte(line))
			if err != nil {
				log.Println("writeToFile err: ", err.Error())
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
}
