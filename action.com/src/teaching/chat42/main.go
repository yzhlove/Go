package main

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"sync"
)

//使用waitgroup优化代码

type Info struct {
	Host string
	Port int
	wg   sync.WaitGroup
	//saveChan chan bytes.Buffer
}

func (info *Info) openSocket(conn net.Conn) {
	defer info.wg.Done()
	var (
		buf     = make([]byte, 1024)
		n       int
		err     error
		byteBuf bytes.Buffer
	)
	for {
		if n, err = conn.Read(buf); err != nil || n == 0 {
			fmt.Printf("[NetErr:%s]\n", err.Error())
			break
		}
		byteBuf.Write(buf)
	}
	fmt.Println("Info = ", byteBuf.String())
	_ = conn.Close()
}

func main() {

	info := &Info{
		Host: "127.0.0.1",
		Port: 1234,
		//saveChan: make(chan bytes.Buffer),
	}
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", info.Host+":"+strconv.Itoa(info.Port)); err != nil {
		panic(err)
	}
	info.wg.Add(1)
	go info.openSocket(conn)
	info.wg.Wait()
	fmt.Printf("Done.")
}
