package worker

import (
	"errors"
	"fmt"
	"journey/chat31_love/engine"
	"journey/chat31_love/zhenai/parser"
	"journey/chat35_distributed/config"
	"log"
)

type SerializedParser struct {
	FnName string
	Args   interface{}
}

type Request struct {
	URL    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		URL: r.URL,
		Parser: SerializedParser{
			FnName: name,
			Args:   args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parse, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		URL:    r.URL,
		Parser: parse,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		parse, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserialize request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, parse)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FnName {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v ", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
