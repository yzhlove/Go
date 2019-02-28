package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, service interface{}) error {
	var (
		err      error
		listener net.Listener
		conn     net.Conn
	)
	if err = rpc.Register(service); err != nil {
		return err
	}
	if listener, err = net.Listen("tcp", host); err != nil {
		return err
	}
	for {
		if conn, err = listener.Accept(); err != nil {
			log.Printf("Accept Err %v ", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", host); err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
