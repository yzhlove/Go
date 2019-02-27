package main

import (
	"fmt"
	rpcdemo "journey/chat34_rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", ":1234"); err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 5}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	fmt.Println(result, err)

}
