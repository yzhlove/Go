package client

import (
	"fmt"
	"journey/chat31_love/engine"
	"journey/chat35_distributed/config"
	"journey/chat35_distributed/rpcsupport"
	"journey/chat35_distributed/worker"
	"net/rpc"
)

func CreateProcessor() (engine.Processor, error) {
	var (
		err    error
		client *rpc.Client
	)
	if client, err = rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0)); err != nil {
		return nil, err
	}
	return func(request engine.Request) (result engine.ParseResult, e error) {
		req := worker.SerializeRequest(request)
		var res worker.ParseResult
		if err := client.Call(config.CrawlServiceRpc, req, &res); err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(res), nil
	}, nil
}
