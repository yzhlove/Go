package tool

import (
	"bufio"
	"net"
)

func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	dataReader := bufio.NewReader(conn)
	for {

	}
}
