package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cf := context.WithCancel(context.TODO())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if cancelThread(ctx) {
					break
				}
				fmt.Println("running:", i)
				time.Sleep(5 * time.Second)
			}
			fmt.Println("[", i, " :Done]")
		}(i, ctx)
	}

	time.Sleep(1 * time.Second)
	cf()
	time.Sleep(6 * time.Second)

}

func cancelThread(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
