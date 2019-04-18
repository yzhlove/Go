package expre

import (
	"expression/chat03/match"
	"log"
	"strings"
)

type Matcher interface {
	Express(key string) error
	Query() error
}

var MatcherList map[string]Matcher

func Run(keys ...string) {
	var (
		matcher Matcher
		err     error
	)
	for _, key := range keys {
		m, ok := MatcherList[strings.ToUpper(strings.Split(key, "_")[0])]
		if ok {
			matcher = m
		} else {
			matcher = match.NewDefaultMatcher()
		}
		if err = matcher.Express(key); err != nil {
			panic("Matcher Err")
		}
		_ = matcher.Query()
	}
}

func RegisterMatcher(match string, matcher Matcher) {
	if _, ok := MatcherList[match]; ok {
		log.Printf("[INFO] %v 已经注册了", match)
	} else {
		MatcherList[match] = matcher
	}
}
