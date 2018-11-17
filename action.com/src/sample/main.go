//Go In Action数据搜索

package main

import (
	"log"
	"os"
	_ "sample/matchers"
	"sample/search"
)

func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	//使用特定项进行搜索
	search.Run("president")
}
