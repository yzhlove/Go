package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "xjj")
	go watch(ctx, "lcm")
	go watch(ctx, "xyj")
	go watch(ctx, "fyb")
	go watch(ctx, "hxy")

	time.Sleep(10 * time.Second)
	fmt.Println("ok")

	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("check")
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出,停止了...")
			return
		case <-time.NewTimer(2 * time.Second).C:
			fmt.Println(name, ":running ... ")
		}
	}
}
