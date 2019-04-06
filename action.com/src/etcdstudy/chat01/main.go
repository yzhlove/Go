package main

import (
	"fmt"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
	)
	// 分钟(0-59) 小时(0-23) 天(1-31) 月(1-12) 星期(0-6)
	if expr, err = cronexpr.Parse("* * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	//每个五分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	expr = expr
}
