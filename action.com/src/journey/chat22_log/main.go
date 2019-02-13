package main

import (
	"journey/chat22_log/model"
	"log"
)

func createLogger() *model.Logger {
	logger := model.NewLogger()
	cw := model.NewConsoleLog()
	logger.Register(cw)

	fw := model.NewFileLogger()
	if err := fw.OptFileLogger("./logger.log"); err != nil {
		log.Println(err.Error())
	} else {
		logger.Register(fw)
	}
	return logger
}

func main() {
	msg := []string{"what are doing", "are you kid me", "are you ok", "what fuck you"}
	logger := createLogger()
	for _, v := range msg {
		logger.WriteLog(v)
	}
}
