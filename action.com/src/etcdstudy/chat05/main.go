package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//etcd GET

func main() {
	var (
		client *clientv3.Client
		err    error
	)
	config := clientv3.Config{
		Endpoints:   []string{":2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Print(err)
		return
	}

	dataEngine := clientv3.NewKV(client)

	//GET
	if getResp, err := dataEngine.Get(context.TODO(), "/cron/jobs/job1" /*clientv3.WithCountOnly()*/); err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Println(getResp.Kvs, " * ", getResp.Count)
	}

}
