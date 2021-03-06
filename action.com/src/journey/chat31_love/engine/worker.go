package engine

import (
	"journey/chat31_love/fetcher"
	"log"
)

func Worker(req Request) (ParseResult, error) {
	var (
		bytes []byte
		err   error
	)
	log.Println(req.URL)
	if bytes, err = fetcher.Fetch(req.URL); err != nil {
		log.Printf("Fetch err url:%s err:%v", req.URL, err)
		return ParseResult{}, err
	}
	return req.Parser.Parse(bytes, req.URL), nil

}
