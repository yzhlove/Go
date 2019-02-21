package main

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/scheduler"
	"journey/chat31_love/zhenai/parser"
)

// 真爱网

const (
	url = "http://www.zhenai.com/zhenghun"
)

func main() {

	/*	engine.SimpleEngine{}.Run(engine.Request{
		URL:       url,
		ParseFunc: parser.ParseCityList,
	})*/

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 10,
	}

	e.Run(engine.Request{
		URL:       url,
		ParseFunc: parser.ParseCityList,
	})

}
