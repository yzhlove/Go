package main

import (
	"fmt"
	"journey/chat31_love/engine"
	"journey/chat31_love/scheduler"
	"journey/chat31_love/zhenai/parser"
	iceservers "journey/chat35_distributed/client"
	"journey/chat35_distributed/config"
	worker "journey/chat35_distributed/worker/client"
)

func main() {
	var (
		itemChan  chan engine.Item
		err       error
		processor engine.Processor
	)
	if itemChan, err = iceservers.ItemSaver(fmt.Sprintf(":%d", config.ItemServerPort)); err != nil {
		panic(err)
	}

	if processor, err = worker.CreateProcessor(); err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueueScheduler{},
		WorkCount:      100,
		ItemChan:       itemChan,
		RequestProcess: processor,
	}
	e.Run(engine.Request{
		URL:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}
