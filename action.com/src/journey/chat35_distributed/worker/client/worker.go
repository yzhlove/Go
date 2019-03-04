package client

import (
	"journey/chat31_love/engine"
	"journey/chat35_distributed/config"
	"journey/chat35_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(request engine.Request) (result engine.ParseResult, e error) {
		req := worker.SerializeRequest(request)
		var res worker.ParseResult
		c := <-clientChan
		if err := c.Call(config.CrawlServiceRpc, req, &res); err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(res), nil
	}
}
