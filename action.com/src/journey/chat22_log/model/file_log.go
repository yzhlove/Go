package model

import (
	"errors"
	"fmt"
	"os"
)

const (
	mode       = os.O_RDWR | os.O_CREATE | os.O_APPEND
	permission = 0666
)

type fileLog struct {
	file    *os.File
	logName string
}

func (f *fileLog) OptFileLogger(path string) error {
	var err error
	if path != "" {
		f.logName = path
	}
	if f.file, err = os.OpenFile(f.logName, mode, permission); err != nil {
		return err
	}
	return nil
}

func (f *fileLog) Info(data interface{}) error {
	if f.file == nil {
		return errors.New("file is nil object")
	}
	info := fmt.Sprintf("%v\n", data)
	_, err := f.file.Write([]byte(info))
	return err
}

func NewFileLogger() *fileLog {
	return &fileLog{
		file:    nil,
		logName: "./development.log",
	}
}
