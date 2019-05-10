package main

import (
	"fmt"
	"journey/chat35_distributed/config"
	"journey/chat35_distributed/rpcsupport"
	"journey/chat35_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		URL: "http://album.zhenai.com/u/1426975040",
		Parser: worker.SerializedParser{
			FnName: config.ParseProfile,
			Args:   "不想",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
