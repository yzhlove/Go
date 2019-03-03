package parser

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/model"
	"journey/chat35_distributed/config"
	"regexp"
)

var userInfoReg = regexp.MustCompile(`<div class="m-btn [a-z]+" data-v-[a-zA-Z0-9]+>([^<]+)</div>`)
var urlReg = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)

func parseProfile(contents []byte, url, name string) engine.ParseResult {
	userProfile := model.Profile{}
	matchers := userInfoReg.FindAllSubmatch(contents, -1)

	userProfile.NickName = name

	if len(matchers) > 10 {
		userProfile.Marriage = string(matchers[0][1])
		userProfile.Age = string(matchers[1][1])
		userProfile.Xinzuo = string(matchers[2][1])
		userProfile.Height = string(matchers[3][1])
		userProfile.Weight = string(matchers[4][1])
		userProfile.Income = string(matchers[6][1])
		userProfile.Occupation = string(matchers[7][1])
		userProfile.Education = string(matchers[8][1])
		userProfile.HoKou = string(matchers[10][1])
	}

	//从URL解析用户的ID
	Id := ""
	results := urlReg.FindStringSubmatch(url)
	if len(results) > 0 {
		Id = results[1]
	}
	return engine.ParseResult{
		Items: []engine.Item{{URL: url, Type: "zhenai", Id: Id, Detail: userProfile}},
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

//func ProfileParser(name string) engine.ParseFunc {
//	return func(bytes []byte, url string) engine.ParseResult {
//		return ParseProfile(bytes, url, name)
//	}
//}
