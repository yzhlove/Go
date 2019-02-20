package main

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/zhenai/parser"
)

// 真爱网

const (
	url = "http://www.zhenai.com/zhenghun"
)

func main() {

	engine.Run(engine.Request{
		URL:       url,
		ParseFunc: parser.ParseCityList,
	})

}
