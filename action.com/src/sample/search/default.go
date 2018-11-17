package search

//默认匹配器
type defaultMatcher struct{}

func init() {
	// Register("default", matcher)
	var matcher defaultMatcher
	Register("default", matcher)
}

//Search 实现默认匹配器行为
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
