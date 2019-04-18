package match

import "expression/chat03/expre"

type DefaultMatcher struct{}

func init() {
	expre.RegisterMatcher("default", &DefaultMatcher{})
}

func (matcher *DefaultMatcher) Query() error {

	return nil
}

func (matcher *DefaultMatcher) Express(key string) error {

	return nil
}

func NewDefaultMatcher() *DefaultMatcher {
	return &DefaultMatcher{}
}
