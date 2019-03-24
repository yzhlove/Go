package tool

import (
	"bufio"
	"net"
)

func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	var (
		pkt Packet
		err error
	)
	dataReader := bufio.NewReader(conn)
	for {
		if pkt, err = readPacket(dataReader); err != nil || !callback(conn, pkt.Body) {
			_ = conn.Close()
			break
		}
	}
}
