package scheduler

import (
	"merge_tools/engine"
	"merge_tools/item"
)

type SimpleScheduler struct {
	RequestChan  chan *item.Item
	ResponseChan chan *item.Item
}

func (s *SimpleScheduler) Request(*engine.Item) {
	panic("implement me")
}

func (s *SimpleScheduler) Response() *engine.Item {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	panic("implement me")
}
