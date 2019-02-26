package engine

type ParseFunc func(contents []byte, url string) ParseResult

type Request struct {
	URL       string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	URL    string
	Id     string
	Type   string
	Detail interface{}
}

func NilParser(b []byte) ParseResult {
	return ParseResult{}
}
