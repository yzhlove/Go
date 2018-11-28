package main

import (
	"bufio"
	"fmt"
	"os"
)

//安行读取文件

const (
	filePath = "/Users/love/logs/nginx_error.log"
)

func readFile() {

	file , err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	readFile()
}