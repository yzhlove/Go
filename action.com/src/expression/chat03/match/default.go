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
	fmt.Println("DEFAULT KEY : ", key)
	return nil
}
