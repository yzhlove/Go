package conn

import (
	"io"
	"log"
	"sync/atomic"
)

var GenerateID int32

type Conn struct {
	ID int32
}

func (c *Conn) Close() error {
	log.Printf("[WRAN] closed connection %v \n", c.ID)
	return nil
}

func New() (io.Closer, error) {
	id := atomic.AddInt32(&GenerateID, 1)
	return &Conn{ID: id}, nil
}
