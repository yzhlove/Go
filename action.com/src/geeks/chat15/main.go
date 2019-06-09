package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// filter模式

func main() {

	sf := NewSplitFilter(",")
	tf := NewToIntFilter()
	smf := NewSumFilter()

	pipe := NewPipeline("pipe", sf, tf, smf)

	if result, err := pipe.Process("1,2,3,4,5"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("result = %T %v \n", result, result)
	}

}

/*----------------- base -----------------*/
type Request interface{}
type Response interface{}
type Filter interface {
	Process(data Request) (Response, error)
}

/*----------------- split filter -----------------*/
type splitFilter struct {
	mode string
}

func NewSplitFilter(data string) *splitFilter {
	return &splitFilter{mode: data}
}

func (sf *splitFilter) Process(req Request) (resp Response, err error) {
	if str, ok := req.(string); ok {
		return strings.Split(str, sf.mode), nil
	}
	return nil, errors.New("type must string")
}

/*----------------- toInt filter -----------------*/
type toInt struct{}

func NewToIntFilter() *toInt {
	return &toInt{}
}

func (tf *toInt) Process(req Request) (resp Response, err error) {
	if args, ok := req.([]string); ok {
		list := make([]int, 0, len(args))
		for _, v := range args {
			if temp, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				list = append(list, temp)
			}
		}
		return list, nil
	}
	return nil, errors.New("type must []string")
}

/*----------------- sum filter -----------------*/
type sumFilter struct{}

func NewSumFilter() *sumFilter {
	return &sumFilter{}
}

func (sf *sumFilter) Process(req Request) (resp Response, err error) {
	if values, ok := req.([]int); ok {
		var sum int
		for _, v := range values {
			sum += v
		}
		return sum, nil
	}
	return nil, errors.New("type must []int")
}

/*----------------- pipe filter -----------------*/
type pipeFilter struct {
	Name    string
	Filters *[]Filter
}

func NewPipeline(name string, filters ...Filter) *pipeFilter {
	return &pipeFilter{
		Name:    name,
		Filters: &filters,
	}
}

func (pf *pipeFilter) Process(req Request) (resp Response, err error) {
	for _, filter := range *pf.Filters {
		if resp, err = filter.Process(req); err != nil {
			return
		}
		req = resp
	}
	return
}
