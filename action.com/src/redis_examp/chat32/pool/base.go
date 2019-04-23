package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	ErrPoolClosed = errors.New("connection must closed")
	ErrPoolGet    = errors.New("time out")
)

type Function func() (io.Closer, error)

type Pool struct {
	mutex    sync.Mutex
	resource chan io.Closer
	closed   bool
	factory  Function
}

func New(factory Function, size uint) (*Pool, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size must > 0")
	}
	p := &Pool{
		factory:  factory,
		resource: make(chan io.Closer, size),
	}
	//初始化连接池
	for i := 0; i < int(size); i++ {
		res, _ := p.factory()
		p.resource <- res
	}
	return p, nil
}

func (p *Pool) Get() (io.Closer, error) {
	select {
	case res, ok := <-p.resource:
		if !ok {
			return nil, ErrPoolClosed
		}
		return res, nil
	default:
		return nil, ErrPoolGet
	}
}

func (p *Pool) Back(res io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		_ = res.Close()
		return
	}
	select {
	case p.resource <- res:
	default:
		_ = res.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resource)
	for r := range p.resource {
		_ = r.Close()
	}
}
