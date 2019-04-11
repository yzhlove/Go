package main

import "fmt"

type Fiture interface {
	Name()
}

type Apple struct{}

func (a *Apple) Name() {

}

func checkFiture() Fiture {
	var app *Apple
	return app
}

func main() {
	f := checkFiture()
	if f == nil {
		fmt.Print("fuck")
	} else {
		fmt.Print("you")
	}
	fmt.Printf("%T %v \n", f, f)

}
