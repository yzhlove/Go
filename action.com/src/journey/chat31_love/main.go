package main

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/persist"
	"journey/chat31_love/scheduler"
	"journey/chat31_love/zhenai/parser"
	"journey/chat35_distributed/config"
)

// 真爱网

const (
	url      = "http://www.zhenai.com/zhenghun"
	shanghai = "http://www.zhenai.com/zhenghun/shanghai"
)

func main() {
	var (
		itemChan chan engine.Item
		err      error
	)
	/*	engine.SimpleEngine{}.Run(engine.Request{
		URL:       url,
		ParseFunc: parser.ParseCityList,
	})*/

	if itemChan, err = persist.ItemSave("dating_profile"); err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueueScheduler{},
		WorkCount:      100,
		ItemChan:       itemChan,
		RequestProcess: engine.Worker,
	}

	//e.Run(engine.Request{
	//	URL:       url,
	//	ParseFunc: parser.ParseCityList,
	//})

	//上海
	e.Run(engine.Request{
		URL: url,
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})

}
