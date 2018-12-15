package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//Pool 缓存池
type Pool struct {
	m        sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

var (
	//ErrPoolClosed 缓存池已经关闭
	ErrPoolClosed = errors.New("Pool has been closed")
)

//New 创建一个缓存池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value tool small")
	}
	return &Pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
	}, nil
}

//Acquire 从缓存池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		log.Println("Acquire:", "Shared Reasource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Reasource")
		return p.factory()
	}
}
