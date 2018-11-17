package search

import (
	"fmt"
	"log"
)

//Result  保存搜索结果
type Result struct {
	Field   string
	Content string
}

//Matcher 搜索接口
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//Match 函数，为每个数据源单独启动一个协程来执行这个函数
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器执行搜索
	searchResults , err := matcher.Search(feed,searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//将数据结果写入通道
	for _,result := range searchResults {
		results <- result
	}
}

//Display 从每个单独的协程接收到结果之后在终端窗口显示
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
