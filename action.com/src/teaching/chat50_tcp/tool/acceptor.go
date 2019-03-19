package tool

import (
	"fmt"
	"net"
	"sync"
)

type Acceptor struct {
	listen net.Listener
	sync.WaitGroup
	OnSessionData func(net.Conn, []byte) bool
}

func (apt *Acceptor) Start(address string) {
	go apt.listening(address)
}

func (apt *Acceptor) listening(address string) {
	apt.Add(1)
	defer apt.Done()

	var (
		err  error
		conn net.Conn
	)
	if apt.listen, err = net.Listen("tcp", address); err != nil {
		fmt.Printf("[ListenErr:%s ]\n", err.Error())
		return
	}
	for {
		if conn, err = apt.listen.Accept(); err != nil {
			fmt.Printf("[AcceptErr:%s ]\n", err.Error())
			break
		}
		go handleSession(conn, apt.OnSessionData)
	}

}

func (apt *Acceptor) Stop() {
	apt.listen.Close()
}

func (apt *Acceptor) Waiter() {
	apt.Wait()
}

func NewAcceptor() *Acceptor {
	return &Acceptor{}
}
