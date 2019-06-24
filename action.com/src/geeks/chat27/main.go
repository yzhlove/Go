package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

//网络测试

func main() {

}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("connect accept connection:%v ", err)
				continue
			}
			connectToService()
			_, _ = fmt.Fprintf(conn, "")
			_ = conn.Close()
		}
	}()
	return &wg
}

func connectToService() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}
