package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//一个日志记录器

var (
	Trace   *log.Logger //记录所有日志
	Info    *log.Logger //重要信息
	Warning *log.Logger //需要注意的信息
	Error   *log.Logger //非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failer to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout, "WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	//io.MultiWriter 会绑定实现write接口的值，并同时想多个write设备写数据
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	Trace.Println("I hava a dream!")
	Info.Println("I hava a dream!")
	Warning.Println("I hava a dream!")
	Error.Println("I hava a dream!")

}
