package main

import (
	"fmt"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept Err:", err.Error())
			continue
		}
		_, _ = conn.Write([]byte("What Are You Doing ..."))
		_ = conn.Close()
	}
}
