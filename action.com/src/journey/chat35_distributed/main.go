package main

import (
	"fmt"
	"journey/chat31_love/engine"
	"journey/chat31_love/scheduler"
	"journey/chat31_love/zhenai/parser"
	iceservers "journey/chat35_distributed/client"
	"journey/chat35_distributed/config"
	"journey/chat35_distributed/rpcsupport"
	worker "journey/chat35_distributed/worker/client"
	"log"
	"net/rpc"
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
	pool := createClientPool([]string{})
	processor = worker.CreateProcessor(pool)

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

func createClientPool(hosts []string) chan *rpc.Client {

	var (
		clients []*rpc.Client
		tempCli *rpc.Client
		err     error
	)
	for _, host := range hosts {
		if tempCli, err = rpcsupport.NewClient(host); err != nil {
			log.Printf("Error connecting %s: %v", host, err)
		} else {
			clients = append(clients, tempCli)
			log.Printf("Connected %s ", host)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, cli := range clients {
				out <- cli
			}
		}
	}()
	return out
}
