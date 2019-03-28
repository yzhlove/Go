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
		listen    net.Listener
		conn      net.Conn
		err       error
		bytes     []byte
		respBytes []byte
		rpcData   RpcData
		funcName  reflect.Value
		ok        bool
	)
	if listen, err = net.Listen("tcp", rpcSev.addr); err != nil {
		panic(err)
	}
	for {
		if conn, err = listen.Accept(); err != nil {
			log.Printf("Accepr Err:%v \n", err)
			continue
		}
		session := NewSession(conn)
		if bytes, err = session.Read(); err != nil {
			log.Printf("SessReadErr:%v \n", err)
			continue
		}
		//解码
		if rpcData, err = decode(bytes); err != nil {
			log.Printf("DecodeErr:%v \n", err)
			continue
		}
		if funcName, ok = rpcSev.funds[rpcData.Name]; !ok {
			log.Printf("Func %s not exists!\n", rpcData.Name)
			continue
		}
		//构造参数
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		//函数调用
		out := funcName.Call(inArgs)
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		//返回数据
		respData := RpcData{Name: rpcData.Name, Args: outArgs}
		if respBytes, err = encode(respData); err != nil {
			log.Printf("EncodeErr:%v \n", err)
			continue
		}

		if err = session.Write(respBytes); err != nil {
			log.Printf("SendErr:%v \n", err)
			continue
		}
		log.Printf("SendSuccessful Data!\n")
	}
}

func (rpcSev *RpcServer) Register(rpcName string, fn interface{}) {
	if _, ok := rpcSev.funds[rpcName]; ok {
		return
	}
	rpcSev.funds[rpcName] = reflect.ValueOf(fn)
}
