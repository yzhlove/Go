package monitor

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

/*
<-chan 只读
chan <- 只写
*/

//Runner 监控
type Runner struct {
	//interrupt通道报告从操作系统
	//发送信号
	interrupt chan os.Signal
	//complete报告处理任务已经完成
	complete chan error
	//timeout报告处理任务已经超时
	timeout <-chan time.Time
	//tasks函数
	tasks []func(int)
}

var (
	//ErrorTimeout 任务执行超时返回
	ErrorTimeout = errors.New("time out")
	//ErrInterrupt 接受到操作系统的事件返回
	ErrInterrupt = errors.New("signal result")
)

//New 返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add 接收一个任务附加到Runner上
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start 开始任务
func (r *Runner) Start() error {
	//接收所有中断的信号
	signal.Notify(r.interrupt, os.Interrupt)

	//不同的goruntine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrorTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测操作系统的中断的信号
		if r.goInterrupt() {
			return ErrInterrupt
		}
		//执行已经注册任务
		log.Printf("id = %d \n", id+1)
		task(id + 1)
	}
	return nil
}

func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		//停止接收后续任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
