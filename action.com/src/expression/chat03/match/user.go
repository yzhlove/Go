package match

import (
	"expression/chat03/expre"
	"regexp"
)

var userRegex = regexp.MustCompile("")

type UserMatcher struct{}

func init() {
	expre.RegisterMatcher("USER", &UserMatcher{})
}

func (matcher *UserMatcher) Express(key string) error {

	return nil
}

func (matcher *UserMatcher) Query() error {

	return nil
}
