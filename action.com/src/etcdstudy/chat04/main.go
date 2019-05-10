package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//向etcd写入数据

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
	//拿到kv存储引擎
	dataEngine := clientv3.NewKV(client)

	//PUT
	if putResp, err := dataEngine.Put(context.TODO(), "/cron/jobs/job1", "running", clientv3.WithPrevKV()); err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Println("Revision:", putResp)
		if putResp.PrevKv != nil {
			fmt.Println("Prevalue:", string(putResp.PrevKv.Value))
		}
	}
}
