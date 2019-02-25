package main

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/persist"
	"journey/chat31_love/scheduler"
	"journey/chat31_love/zhenai/parser"
)

// 真爱网

const (
	url      = "http://www.zhenai.com/zhenghun"
	shanghai = "http://www.zhenai.com/zhenghun/shanghai"
)

func main() {

	/*	engine.SimpleEngine{}.Run(engine.Request{
		URL:       url,
		ParseFunc: parser.ParseCityList,
	})*/

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
		ItemChan:  persist.ItemSave(),
	}

	//e.Run(engine.Request{
	//	URL:       url,
	//	ParseFunc: parser.ParseCityList,
	//})

	//上海
	e.Run(engine.Request{
		URL:       shanghai,
		ParseFunc: parser.ParseCity,
	})

}
