package main

import (
	"etcdstudy/chat02/work"
	"time"

	"fmt"
)

var (
	scheduleTable map[string]*work.CronTab
	shell         = "*/5 * * * * * *"
)

func main() {
	scheduleTable = make(map[string]*work.CronTab)

	scheduleTable["job1"] = work.NewCrontab(shell)
	scheduleTable["job2"] = work.NewCrontab(shell)

	go checkWork()

	time.Sleep(time.Minute * 5)

}

func checkWork() {
	var (
		jobName string
		crontab *work.CronTab
	)
	for {
		nowTime := time.Now()
		for jobName, crontab = range scheduleTable {
			//判断是否过期
			if crontab.NextTime.Before(nowTime) || crontab.NextTime.Equal(nowTime) {
				go func(jonName string) {
					//启动协程，执行任务
					fmt.Printf("[WorkStart: %v ]\n", jobName)
				}(jobName)
				//更新下一次的执行时间
				crontab.NextTime = crontab.Expr.Next(nowTime)
				fmt.Printf("[NextTime:%v %v]\n", jobName, crontab.NextTime)
			}
		}

		select {
		case <-time.NewTimer(100 * time.Millisecond).C:
		}
	}
}
