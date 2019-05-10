package main

import "fmt"

type Pay interface {
	Swipe()
}

type Ali struct{}

func (Ali) Swipe() {
	fmt.Println("use ali pay")
}

type BankPay struct {
	Ali
}

func (BankPay) Swipe() {
	fmt.Println("use bank pay")
}

func pay(p Pay) {
	p.Swipe()
}

func main() {
	pay(new(BankPay))
	pay(new(Ali))

	var bank = BankPay{}
	pay(bank)
	pay(bank.Ali)

	var ali = Ali{}
	pay(ali)

}
