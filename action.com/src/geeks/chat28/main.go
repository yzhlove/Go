package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {

}

func connectToServer() interface{} {
	time.Sleep(time.Millisecond * 100)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {

	p := &sync.Pool{
		New: connectToServer,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}
func startNetwork() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("connect listen err:%v ", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("connect accept err")
				continue
			}
			svcConn := connPool.Get()
			_, _ = fmt.Fprintf(conn, " 1")
			connPool.Put(svcConn)
			_ = conn.Close()
		}
	}()
	return &wg
}
