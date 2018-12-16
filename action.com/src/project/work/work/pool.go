package work

import "sync"

//Worker 工作池
type Worker interface {
	Task()
}

//Pool 协程 池
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//New 创建一个协程池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//Run 提交任务到协程池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//Shutdown 所有的协程停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
