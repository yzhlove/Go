package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func checkFileExists(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func saveFile(fileName string, fileData []byte) bool {

	var file *os.File
	var err error

	if checkFileExists(fileName) {
		file, err = os.OpenFile(fileName, os.O_RDWR, 0666)
	} else {
		file, err = os.Create(fileName)
	}
	if nil != err {
		log.Fatalln(err)
		return false
	}
	defer file.Close()
	file.Write(fileData)
	return true
}

func getDataByURL(url string) ([]byte, error) {

	response, err := http.Get(url)
	if nil != err {
		log.Fatalln("DownERR:", err.Error())
		return nil, errors.New("DownERR")
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if nil != err {
		log.Fatalln("read Response ERR:", err.Error())
		return nil, errors.New("ReadERR")
	}
	return responseBody, nil
}

func main() {

	listURL := map[string]string{
		"bing":  "https://cn.bing.com/",
		"baidu": "http://www.baidu.com/",
	}

	for name, url := range listURL {
		responseData, err := getDataByURL(url)
		if nil != err {
			log.Fatalln(err.Error())
			break
		}
		saveFile(name, responseData)
	}

}
