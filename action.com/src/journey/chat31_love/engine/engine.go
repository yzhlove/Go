package engine

import (
	"journey/chat31_love/fetcher"
	"log"
)

func Run(seeds ...Request) {
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
		log.Println(req.URL)
		if body, err = fetcher.Fetch(req.URL); err != nil {
			log.Printf("Fetch err url:%s err:%v", req.URL, err)
			continue
		}

		parseRes = req.ParseFunc(body)
		requests = append(requests, parseRes.Requests...)

		for _, item := range parseRes.Items {
			log.Printf("Get Item %v ", item)
		}
	}

}
