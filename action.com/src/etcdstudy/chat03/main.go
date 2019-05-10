package main

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//etcd test
func main() {
	config := clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	}
	var (
		client *clientv3.Client
		err    error
	)
	if client, err = clientv3.New(config); err != nil {
		fmt.Print(err)
		return
	}
	client = client
}
