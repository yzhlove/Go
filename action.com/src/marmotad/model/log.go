package model

import (
	"github.com/davyxu/golog"
)

var log = golog.New("登陆服务")

func init() {
	golog.SetLevelByString("登陆服务", "debug")
}
