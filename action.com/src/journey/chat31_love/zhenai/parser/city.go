package parser

import (
	"journey/chat31_love/engine"
	"regexp"
)

var cityReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
var moreReg = regexp.MustCompile(`"(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {

	matchers := cityReg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, match := range matchers {
		url := string(match[1])
		nickname := string(match[2])
		//result.Items = append(result.Items, "User:"+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL: url,
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, url, nickname)
			},
		})
	}

	matchers = moreReg.FindAllSubmatch(contents, -1)
	for _, match := range matchers {
		//log.Printf("MoreURL -> %s ", string(match[1]))
		result.Requests = append(result.Requests, engine.Request{
			URL:       string(match[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
