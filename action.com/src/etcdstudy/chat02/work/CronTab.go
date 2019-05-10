package work

import (
	"time"

	"github.com/gorhill/cronexpr"
)

type CronTab struct {
	Expr     *cronexpr.Expression
	NextTime time.Time
}

func NewCrontab(shell string) *CronTab {
	return &CronTab{
		Expr: cronexpr.MustParse(shell),
	}
}
