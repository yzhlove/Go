package parser

import (
	"journey/chat31_love/engine"
	"journey/chat35_distributed/config"
	"regexp"
)

const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {

	reg := regexp.MustCompile(cityListRegex)
	matchers := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range matchers {
		result.Requests = append(result.Requests, engine.Request{
			URL: string(match[1]),
			Parser: engine.NewFuncParser(ParseCity,
				config.ParseCity),
		})
	}
	return result
}
