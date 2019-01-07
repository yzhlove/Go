package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	const filePath = "/Users/love/WorkSpace/Go/action.com/src/basic/chat05/config.ini"

	search := make(map[string]string)
	search[`remote "origin"`] = "fetch"
	search["core"] = "hideDotFiles"

	for k, v := range search {
		result := getConfigFile(filePath, k, v)
		fmt.Printf("%v + %v = %v \n", k, v, result)
	}
}

func getConfigFile(filename, expectSection, exceptKey string) string {

	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var temp string
	for {
		var liner string
		if liner, err = reader.ReadString('\n'); err != nil {
			break
		}
		liner = strings.TrimSpace(liner)
		if liner == "" || liner[0] == ';' {
			continue
		}
		if liner[0] == '[' && liner[len(liner)-1] == ']' {
			temp = liner[1 : len(liner)-1]
		} else if temp == expectSection {
			pair := strings.Split(liner, "=")
			if len(pair) == 2 {
				if strings.TrimSpace(pair[0]) == exceptKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}
