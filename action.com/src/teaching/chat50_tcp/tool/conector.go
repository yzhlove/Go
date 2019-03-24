package tool

import (
	"fmt"
	"net"
	"strconv"
)

func Connector(address string, sendTimes int) {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", address); err != nil {
		fmt.Printf("tcp conn err:%v \n", err)
		return
	}
	for i := 0; i < sendTimes; i++ {
		tmpStr := strconv.Itoa(i)
		if err = writePacket(conn, []byte(tmpStr)); err != nil {
			fmt.Printf("send packet err:%v \n", err)
			break
		}
	}
}
