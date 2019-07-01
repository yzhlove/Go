package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

//cron解析

func main() {
	var (
		expr *cronexpr.Expression
		err  error
		now  time.Time
		next time.Time
	)
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		panic(err)
	}

	now = time.Now()
	next = expr.Next(now)
	time.AfterFunc(next.Sub(now), func() {
		fmt.Println("ok")
	})
	time.Sleep(time.Second * 6)
	fmt.Println("Done .")
}
