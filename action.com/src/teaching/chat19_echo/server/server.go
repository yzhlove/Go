package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//Echo回声服务器

type EchoServer struct {
	Addr     string
	ExitChan chan bool
}

func (es *EchoServer) Start() {
	var (
		listen net.Listener
		err    error
		conn   net.Conn
	)
	if listen, err = net.Listen("tcp", es.Addr); err != nil {
		fmt.Println(err.Error())
		es.ExitChan <- true
	}
	fmt.Println("[LISTEN] " + es.Addr)
	defer listen.Close()
	for {
		if conn, err = listen.Accept(); err != nil {
			fmt.Println("[Accept Error] ", err.Error())
			continue
		}
		go es.HandleSession(conn)
	}
}

func (es *EchoServer) HandleSession(conn net.Conn) {
	var (
		err    error
		reader *bufio.Reader
		data   string
	)
	fmt.Println("[ -- session start -- ]")
	reader = bufio.NewReader(conn)
	for {
		if data, err = reader.ReadString('\n'); err != nil {
			fmt.Printf("[ --  session closed -- ] [ERROR:%s] \n", err.Error())
			_ = conn.Close()
			break
		}
		data = strings.TrimSpace(data)
		if !es.TelnetCmd(data) {
			_ = conn.Close()
			break
		}
		_, _ = conn.Write([]byte(data + "\r\n"))
	}
}

func (es *EchoServer) TelnetCmd(str string) bool {
	switch {
	case strings.HasPrefix(str, "@close"):
		fmt.Println("[ -- session close -- ]")
		return false
	case strings.HasPrefix(str, "@shutdown"):
		fmt.Println("[ -- server close --]")
		es.ExitChan <- true
		return false
	default:
		fmt.Printf("[Accept Data: %s]\n", str)
		return true
	}
}

func (es *EchoServer) StartLoop() {
	<-es.ExitChan
	fmt.Println("[ -- Done. -- ]")
	os.Exit(0)
}
