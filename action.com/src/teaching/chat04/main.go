package main

import "fmt"

//接口断言

type PayFace interface {
	CanFaceID()
}

type PayQR interface {
	CanQRID()
}

type Alipay struct{}

func (Alipay) CanFaceID() {
	fmt.Println("ali face++ ")
}

type Wechatpay struct{}

func (Wechatpay) CanQRID() {
	fmt.Println("whchat qr code")
}

func Pay(pay interface{}) {
	switch pay.(type) {
	case PayFace:
		pay.(PayFace).CanFaceID()
	case PayQR:
		pay.(PayQR).CanQRID()
	default:
		panic("unknown interface")
	}
}

func main() {

	Pay(new(Alipay))
	Pay(new(Wechatpay))
	Pay("Hello World")

}
