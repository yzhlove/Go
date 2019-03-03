package engine

import (
	"fmt"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var (
		requests []Request
		body     []byte
		err      error
		parseRes ParseResult
	)
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		if parseRes, err = worker(req); err != nil {
			continue
		}
		parseRes = req.Parser.Parse(body, req.URL)
		requests = append(requests, parseRes.Requests...)
		for _, item := range parseRes.Items {
			fmt.Printf("Get Item %v \n", item)
		}
	}

}
