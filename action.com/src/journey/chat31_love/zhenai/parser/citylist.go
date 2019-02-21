package parser

import (
	"journey/chat31_love/engine"
	"regexp"
)

const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(cityListRegex)
	matchers := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range matchers {
		result.Items = append(result.Items, "City:"+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:       string(match[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
