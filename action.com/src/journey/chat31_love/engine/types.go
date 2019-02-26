package engine

type Request struct {
	URL       string
	ParseFunc func([]byte) ParseResult
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
