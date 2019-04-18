package match

import (
	"expression/chat03/expre"
	"fmt"
	"regexp"
)

var userRegex = regexp.MustCompile(`User_[a-zA-Z]+:[\{]?([A-Za-z0-9]+)[\}]?[:]?`)

type UserMatcher struct{}

func init() {
	expre.RegisterMatcher(expre.User, &UserMatcher{})
}

func (matcher *UserMatcher) Express(key string) error {

	if !userRegex.MatchString(key) {
		return fmt.Errorf("[ERROR] 没有匹配到 %v \n", key)
	}
	keyList := userRegex.FindAllString(key, -1)
	fmt.Println(keyList)
	return nil
}
