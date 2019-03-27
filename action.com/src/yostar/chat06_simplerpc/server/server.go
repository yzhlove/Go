package server

import (
	"log"
	"net"
	"reflect"
)

type RpcServer struct {
	addr  string
	funds map[string]reflect.Value
}

func NewRpcServer(address string) *RpcServer {
	return &RpcServer{addr: address, funds: make(map[string]reflect.Value)}
}

func (rpcSev *RpcServer) Run() {
	var (
		listen net.Listener
		conn net.Conn
		err error
	)
	if listen , err = net.Listen("tcp",rpcSev.addr);err != nil {
		panic(err)
	}
	for {
		if conn , err = listen.Accept();err != nil {
			log.Printf("Accepr Err:%v \n",err)
			continue
		}
		sevSession :=
	}
}
