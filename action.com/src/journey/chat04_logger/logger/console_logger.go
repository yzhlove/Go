package logger

import (
	"fmt"
	"os"
)

type consoleWrite struct{}

func (console *consoleWrite) Write(data interface{}) error {
	str := fmt.Sprintf("%v", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func NewConsoleWrite() *consoleWrite {
	return &consoleWrite{}
}
