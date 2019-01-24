package main

import "fmt"

//结构体内嵌

//BasicColor 基本颜色
type BasicColor struct {
	R int
	G int
	B int
}

type StandColor struct {
	BasicColor //内嵌结构体
	alpha      float64
}

func main() {

	color := new(StandColor)
	color.G = 3
	color.R = 3
	color.B = 3
	color.alpha = 0.5

	fmt.Println("color:", color)
}
