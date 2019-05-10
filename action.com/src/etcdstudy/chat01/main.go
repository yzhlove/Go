package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		nowTime  time.Time
		nextTime time.Time
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

	//下一次的调度时间
	//注意:下一次的调度时间是按(假如每五分钟调度一次)0，5，10，15，20 ...
	//而不是按当前时间计算，
	nowTime = time.Now()
	nextTime = expr.Next(nowTime)
	fmt.Printf("NowTime = %v NextTime=%v \n", nowTime, nextTime)

	time.AfterFunc(nextTime.Sub(nowTime), func() {
		fmt.Println("调度了")
	})
	time.Sleep(time.Second * 120)
}
