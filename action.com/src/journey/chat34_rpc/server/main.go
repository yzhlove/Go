package main

import (
	rpcdemo "journey/chat34_rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	var (
		listener net.Listener
		err      error
		conn     net.Conn
	)
	_ = rpc.Register(rpcdemo.DemoService{})
	if listener, err = net.Listen("tcp", ":1234"); err != nil {
		panic(err)
	}

	for {
		if conn, err = listener.Accept(); err != nil {
			log.Printf("accept err : %v ", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}
