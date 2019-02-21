package parser

import (
	"journey/chat31_love/engine"
	"regexp"
)

const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(cityReg)
	matchers := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, match := range matchers {
		nickname := string(match[2])
		result.Items = append(result.Items, "User:"+nickname)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(match[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, nickname)
			},
		})
	}
	return result
}
