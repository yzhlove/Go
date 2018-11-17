package search

import (
	"log"
	"sync"
)

//注册用于搜索的匹配器
var matchers = make(map[string]Matcher)

//Run 执行搜索逻辑
func Run(searchTerm string) {
	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//创建一个无缓冲通道，接受匹配之后的结果
	results := make(chan *Result)

	//构造一个waitGroup，处理所有额数据源
	var waitGroup sync.WaitGroup

	//设置需要处理数据源的协程的数量
	waitGroup.Add(len(feeds))

	//为每个数据源启动一个协程
	for _, feed := range feeds {
		//获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//启动一个协程执行搜索
		go func(matcher Matcher , feed *Feed) {
			Match(matcher,feed,searchTerm,results)
			waitGroup.Done()
		}(matcher,feed)

	}

	//启动一个协程监控所有的工作是否做完
	go func() {
		waitGroup.Wait()
		//关闭通道
		close(results)
	}()

	Display(results)

}

//Register 注册器
func Register(freeType string, matcher Matcher) {
	if _, exists := matchers[freeType]; exists {
		log.Fatalln(freeType, "Matcher already registered!")
	}
	log.Println("Register", freeType, "matcher")
	matchers[freeType] = matcher
}
