package main

import (
	"net"
	"strconv"
	"teaching/chat50_tcp/tool"
)

//tcp粘包处理

const (
	TestIndex = 100
	address   = ":1234"
)

func main() {
	var recvIndex int
	accept := tool.NewAcceptor()
	accept.Start(address)
	accept.OnSessionData = func(conn net.Conn, bytes []byte) bool {
		tmpStr := string(bytes)
		toInt, err := strconv.Atoi(tmpStr)
		if err != nil || recvIndex != toInt {
			panic("failed")
		}
		recvIndex++
		if recvIndex >= TestIndex {
			accept.Stop()
			return false
		}
		return true
	}
	tool.Connector(address, TestIndex)
	accept.Wait()
}
