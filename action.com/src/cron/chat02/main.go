package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorhill/cronexpr"
)

//多任务调度

var (
	schedTable map[int]*Task
	counter    int
)

type Task struct {
	Next time.Time
	Expr *cronexpr.Expression
}

func init() {
	schedTable = make(map[int]*Task)
}

func main() {

	temp1, err := NewTask("*/5 * * * * * *")
	if err != nil {
		panic(err)
	}

	temp2, err := NewTask("*/7 * * * * * *")
	if err != nil {
		panic(err)
	}

	AddTask(temp1)
	AddTask(temp2)

	go func() {
		for {
			now := time.Now()
			for id, task := range schedTable {
				if task.Next.Before(now) || task.Next.Equal(now) {
					go func(id int) {
						fmt.Println(id, ":执行了...")
					}(id)
					task.Next = task.Expr.Next(now)
					fmt.Println("NextTime = ", task.Next)
				}
			}
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(100 * time.Second)

}

func NewTask(crontab string) (*Task, error) {
	expr, err := cronexpr.Parse(crontab)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &Task{Expr: expr, Next: expr.Next(now)}, nil
}

func AddTask(task *Task) int {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	schedTable[counter] = task
	return counter
}
