package main

import "log"

//

func init() {
	//设置日志前缀
	log.SetPrefix("TRACE: ")
	//日期 | 毫秒级时间戳 | 文件名元素和行号
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	log.Println("message")

	log.Fatalln("fatal message")

	log.Panicln("panic message")

}
