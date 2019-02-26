package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()
	//e.Scheduler.ConfigureMasterWorkChan(in)
	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.URL) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			if result, err := worker(request); err != nil {
				continue
			} else {
				out <- result
			}
		}
	}()
}

var mapeURL = make(map[string]bool)

func isDuplicate(url string) bool {
	if v, ok := mapeURL[url]; ok && v {
		return true
	}
	mapeURL[url] = true
	return false
}
