package model

import (
	"fmt"
	"os"
)

type consoleLog struct{}

func (cl *consoleLog) Info(data interface{}) error {
	logInfo := fmt.Sprintf("%v", data)
	_, err := os.Stdout.Write([]byte(logInfo))
	return err
}

func NewConsoleLog() *consoleLog {
	return &consoleLog{}
}
