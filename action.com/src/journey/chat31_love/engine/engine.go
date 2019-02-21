package engine

import (
	"fmt"
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
		if parseRes, err = worker(req); err != nil {
			continue
		}
		parseRes = req.ParseFunc(body)
		requests = append(requests, parseRes.Requests...)
		for _, item := range parseRes.Items {
			fmt.Printf("Get Item %v \n", item)
		}
	}

}

func worker(req Request) (ParseResult, error) {
	var (
		bytes []byte
		err   error
	)
	log.Println(req.URL)
	if bytes, err = fetcher.Fetch(req.URL); err != nil {
		log.Printf("Fetch err url:%s err:%v", req.URL, err)
		return ParseResult{}, err
	}
	return req.ParseFunc(bytes), nil

}
