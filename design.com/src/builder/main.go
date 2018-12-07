package main

import (
	"errors"
	"fmt"
)

//构建模式

//定义灯的基本信息

//LightColor 灯的颜色
type LightColor string

//LightStatus 灯的状态
type LightStatus bool

//LightBrand 灯的品牌
type LightBrand string

//定义等的颜色常量
const (
	//RedColor 红色
	RedColor LightColor = "red"
	//GreenColor 绿色
	GreenColor = "green"
	//BlueColor 蓝色
	BlueColor = "blue"
)

//定义品牌的常量
const (
	AppleLight LightBrand = "apple"
	SonyLight             = "sony"
	NokiaLight            = "nokia"
)

//Light 灯的结构定义
type Light struct {
	color  LightColor
	brand  LightBrand
	status LightStatus
}

//LightBuilder 生产灯
type LightBuilder struct {
	Light
}

//LightOption 灯的相关操作
type LightOption interface {
	Open() error
	Close() error
	LightInfo()
}

//=================== 灯的信息设置 ===================

//SetColor 设置等的颜色
func (light LightBuilder) SetColor(color LightColor) LightBuilder {
	light.color = color
	return light
}

//SetBrand 设置灯品牌
func (light LightBuilder) SetBrand(brand LightBrand) LightBuilder {
	light.brand = brand
	return light
}

//=================== 灯的相关操作 ===================

//Open 开灯
func (light Light) Open() error {
	if light.status {
		return errors.New("light status err!")
	}
	fmt.Println("Open Light!")
	light.status = true
	return nil
}

//Close 关灯
func (light Light) Close() error {
	if !light.status {
		return errors.New("light status err!")
	}
	fmt.Println("Close Light!")
	light.status = false
	return nil
}

//LightInfo 灯的信息
func (light Light) LightInfo() {
	fmt.Printf("LightInfo: %+v \n", light)
}

//Build 灯
func (light LightBuilder) Build() LightOption {
	lightObj := Light{
		color:  light.color,
		brand:  light.brand,
		status: false,
	}
	return lightObj
}

//NewBuilder 生产灯
func NewBuilder() LightBuilder {
	return LightBuilder{}
}

func main() {

	build := NewBuilder()
	light := build.SetColor(GreenColor).SetBrand(AppleLight).Build()
	light.Open()
	light.LightInfo()

	light = build.SetColor(RedColor).SetBrand(SonyLight).Build()
	light.Close()
	light.LightInfo()

}
