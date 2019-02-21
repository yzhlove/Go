package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkChan(in)
	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Get Item:#%d %v \n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult) {

	go func() {
		for {
			request := <-in
			if result, err := worker(request); err != nil {
				continue
			} else {
				out <- result
			}
		}
	}()

}
