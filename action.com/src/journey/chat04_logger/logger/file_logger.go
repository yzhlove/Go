package logger

import (
	"errors"
	"fmt"
	"os"
)

type fileWrite struct {
	file *os.File
}

//SetFile 设置文件名
func (f *fileWrite) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(filename)
	if err != nil {
		fmt.Println("FileErr:", err.Error())
	}
	return err
}

//Write 实现write接口
func (f *fileWrite) Write(data interface{}) error {

	if f.file != nil {
		return errors.New("file not create")
	}
	str := fmt.Sprintf("%v", data)
	_, err := f.file.Write([]byte(str))
	return err
}

//NewFileWrite 创建一个fileWrite
func NewFileWrite() *fileWrite {
	return &fileWrite{}
}
