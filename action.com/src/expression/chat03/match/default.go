package match

import (
	"expression/chat03/expre"
	"fmt"
)

type DefaultMatcher struct{}

func init() {
	expre.RegisterMatcher(expre.Default, &DefaultMatcher{})
}

func (matcher *DefaultMatcher) Express(key string) error {
	return fmt.Errorf("[ERROR][DEFAULT] 没有匹配到 %v \n", key)
}
