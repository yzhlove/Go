package scheduler

import "journey/chat31_love/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(request chan engine.Request) {
	s.workerChan = request
}
