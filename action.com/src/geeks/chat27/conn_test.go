package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	daemonStart := startNetworkDaemon()
	daemonStart.Wait()
}

func Benchmark_Network(t *testing.B) {
	for i := 0; i < t.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			t.Error(err)
			return
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			t.Error(err)
			return
		}
		_ = conn.Close()
	}
}
