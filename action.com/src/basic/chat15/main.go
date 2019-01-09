package main

import (
	"fmt"
	"os"
)

//关于空类型的思考

type FileEngine interface {
	Open()
}

type ConfigFile struct {
	file *os.File
	path string
}

//解决方案
//func New(path string) FileEngine

func New(path string) *ConfigFile {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	return &ConfigFile{file: file, path: path}
}

func (c *ConfigFile) Open() {
	defer func() {
		if err := c.file.Close(); err != nil {
			fmt.Println("Close File Err!")
		}
	}()
	fmt.Println("path = " + c.path)
}

func main() {

	// nil 与 interface{} 不相等

	//解决方案
	var fileEngine FileEngine
	var tempFileEngine *ConfigFile
	tempFileEngine = New("./errors.txt")
	if tempFileEngine != nil {
		fileEngine = tempFileEngine
	} else {
		fmt.Printf("New Error")
		return
	}
	fileEngine.Open()
}
