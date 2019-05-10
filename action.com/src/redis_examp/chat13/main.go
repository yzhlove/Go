package main

import (
	"context"
	"time"

	"fmt"
)

//context使用

func main() {

	ctx, cancle := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止")
				return
			case <-time.NewTimer(time.Second).C:
				fmt.Println("正在监控")
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("Stop...")
	cancle()
	time.Sleep(time.Second)
	fmt.Println("End...")
}
