package match

import (
	"expression/chat03/expre"
	"fmt"
	"regexp"
)

var gameRegex = regexp.MustCompile(`Game_[a-zA-Z]+:[0-9]+`)

type GameMatcher struct{}

func init() {
	expre.RegisterMatcher(expre.Game, &GameMatcher{})
}

func (matcher *GameMatcher) Express(key string) error {
	if !gameRegex.MatchString(key) {
		return fmt.Errorf("[ERROR] 没有匹配到 %v \n", key)
	}
	keyList := gameRegex.FindAllString(key, -1)
	fmt.Println(keyList)
	return nil
}
