package parser

import (
	"fmt"
	"journey/chat31_love/engine"
	"regexp"
)

const userReg = `<div class="m-btn purple" data-v-[a-zA-Z0-9]+>([\d]+)岁</div>`

func ParseProfile(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(userReg)
	matchers := reg.FindAllSubmatch(contents, -1)

	for _, match := range matchers {
		fmt.Printf("%s \n", match)
	}

	return engine.ParseResult{}
}
