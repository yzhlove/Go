package main

import (
	"fmt"
	"journey/chat04_logger/logger"
)

func createLogger() *logger.Logger {
	l := logger.NewLogger()
	cw := logger.NewConsoleWrite()
	l.Register(cw)
	fw := logger.NewFileWrite()
	if err := fw.SetFile("development.log"); err != nil {
		fmt.Println(err)
	}
	l.Register(fw)
	return l
}

func main() {

	l := createLogger()
	l.Log("what are you doing")

}
