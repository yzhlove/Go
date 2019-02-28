package main

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/model"
	"journey/chat35_distributed/rpcsupport"
	"net/rpc"
	"testing"
	"time"
)

func TestItemSave(t *testing.T) {
	var (
		client *rpc.Client
		err    error
		host   = ":1234"
	)

	go ServeRpc(host, "test")
	time.Sleep(time.Second)
	if client, err = rpcsupport.NewClient(host); err != nil {
		panic(err)
	}

	excepted := engine.Item{
		URL:  "http://album.zhenai.com/u/1769167712",
		Id:   "1769167712",
		Type: "zhenai",
		Detail: model.Profile{
			Age:        "35岁",
			Height:     "168",
			Weight:     "57",
			Income:     "3000-5000元",
			Gender:     "女",
			NickName:   "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	result := ""
	if err = client.Call("ItemSaveService.Save", excepted, &result); err != nil {
		t.Errorf("save err %v ", err)
	}
	t.Logf("rpc save successful  = %v", result)
}
