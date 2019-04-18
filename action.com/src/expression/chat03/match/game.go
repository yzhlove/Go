package match

import (
	"expression/chat03/expre"
	"regexp"
)

var gameRegex = regexp.MustCompile("")

type GameMatcher struct{}

func (matcher *GameMatcher) Query() error {

	return nil
}

func init() {
	expre.RegisterMatcher("game", &GameMatcher{})
}
