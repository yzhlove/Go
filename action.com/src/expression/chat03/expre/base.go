package expre

import (
	"log"
	"strings"
)

type Matcher interface {
	Express(key string) error
}

const (
	Default = "DEFAULT"
	Game    = "GAME"
	User    = "USER"
)

var MatcherList = make(map[string]Matcher)

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
			matcher = MatcherList[Default]
		}
		if err = matcher.Express(key); err != nil {
			log.Println(err)
		}
	}
}

func RegisterMatcher(match string, matcher Matcher) {
	if _, ok := MatcherList[match]; ok {
		log.Printf("[INFO] %v 已经注册了", match)
	} else {
		MatcherList[match] = matcher
	}
}
