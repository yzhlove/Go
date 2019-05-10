package main

import (
	"teaching/chat19_echo/server"
)

func main() {

	es := &server.EchoServer{
		Addr:     ":1234",
		ExitChan: make(chan bool),
	}
	go es.Start()
	es.StartLoop()
}
