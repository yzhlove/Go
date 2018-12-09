package common

import (
	"fmt"
)

var (
	mode string
	//RunMode 游戏模式
	RunMode = []string{"debug", "online"}
	//Code 渠道
	Code = "yuewan"
	//SECRET 密钥
	SECRET = "ca414f5145129530e061faaeb55c9dc8"
)

//SetMode 设置游戏模式
func SetMode(m string) {
	mode = m
}

//GetMode 获取游戏模式
func GetMode() string {
	return mode
}

//SetCode 设置游戏Code
func SetCode(c string) {
	Code = c
}

//GetCode 获取游戏Code
func GetCode() string {
	return Code
}

//GetLoginAddr 获取登陆地址
func GetLoginAddr(requestURL string) string {
	return fmt.Sprintf("%s/game/game_histories/detail", requestURL)
}
