package scheduler

import (
	"journey/chat31_love/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (qs *QueueScheduler) Submit(req engine.Request) {
	qs.requestChan <- req
}

func (qs *QueueScheduler) WorkerReady(worker chan engine.Request) {
	qs.workerChan <- worker
}

func (qs *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (qs *QueueScheduler) Run() {
	qs.workerChan = make(chan chan engine.Request)
	qs.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activReq engine.Request
			var activWork chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activReq = requestQ[0]
				activWork = workerQ[0]
			}
			select {
			case r := <-qs.requestChan:
				requestQ = append(requestQ, r)
			case w := <-qs.workerChan:
				workerQ = append(workerQ, w)
			case activWork <- activReq:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
