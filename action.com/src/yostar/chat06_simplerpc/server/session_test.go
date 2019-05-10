package server

import (
	"net"
	"sync"
	"testing"
)

const address = "0:2333"
const msg = "hello world"

func Test_SessionRead(t *testing.T) {
	var (
		wg     sync.WaitGroup
		listen net.Listener
		conn   net.Conn
		err    error
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		if listen, err = net.Listen("tcp", address); err != nil {
			t.Fatalf("Listen Err:%v ", err)
		}
		if conn, err = listen.Accept(); err != nil {
			t.Fatalf("Accept Error:%v ", err)
		}
		sess := Session{conn: conn}
		if err = sess.Write([]byte(msg)); err != nil {
			t.Fatalf("Write Err:%v", err)
		}
	}()
	go func() {
		defer wg.Done()
		var (
			conn net.Conn
			err  error
			data []byte
		)
		if conn, err = net.Dial("tcp", address); err != nil {
			t.Fatalf("Client Err:%v ", err)
		}
		sess := Session{conn: conn}
		if data, err = sess.Read(); err != nil {
			t.Fatalf("Read Err:%v ", err)
		}
		t.Logf("Read Message:%v", string(data))
	}()
	wg.Wait()
}
