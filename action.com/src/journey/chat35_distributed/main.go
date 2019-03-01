package main

import (
	"fmt"
	"journey/chat31_love/engine"
	"journey/chat31_love/scheduler"
	"journey/chat31_love/zhenai/parser"
	"journey/chat35_distributed/client"
	"journey/chat35_distributed/config"
)

func main() {
	var (
		itemChan chan engine.Item
		err      error
	)
	if itemChan, err = client.ItemSaver(fmt.Sprintf(":%d", config.ItemServerPort)); err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	})
}
