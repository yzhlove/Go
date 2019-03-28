package server

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) callRpc(rpcName string, funcPrt interface{}) {
	fn := reflect.ValueOf(funcPrt).Elem()

	f := func(args []reflect.Value) []reflect.Value {
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		cliSession := NewSession(c.conn)
		reqRpc := RpcData{Name: rpcName, Args: inArgs}
		bytes, err := encode(reqRpc)
		if err != nil {
			panic(err)
		}
		err = cliSession.Write(bytes)
		if err != nil {
			panic(err)
		}
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		respRpc, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		outArgs := make([]reflect.Value, 0, len(respRpc.Args))
		for i, arg := range respRpc.Args {
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
