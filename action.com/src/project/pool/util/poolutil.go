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

//Release 将一个使用后的资源放回pool
func (p *Pool) Release(r io.Closer) {
	//保证操作安全
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resource <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release", "Closing")
		r.Close()
	}
}

//Close pool停止工作，并关闭现有资源
func (p *Pool) Close() {
	//保证本次操作安全
	p.m.Lock()
	defer p.m.Unlock()
	//检查pool是否关闭
	if p.closed {
		return
	}
	//关闭pool
	p.closed = true
	//在清空通道里的资源之前，需要将通道先关闭
	//如果不这样做，将会发生死锁
	close(p.resource)

	//关闭资源
	for r := range p.resource {
		r.Close()
	}
}
