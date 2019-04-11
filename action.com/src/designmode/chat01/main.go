package main

import "fmt"

//普通工厂

type Option interface {
	Operator(int, int) int
}

type Add struct{}

func (add *Add) Operator(a, b int) int {
	return a + b
}

type Multiple struct{}

func (multiple *Multiple) Operator(a, b int) int {
	return a * b
}

type OptionFactory struct{}

func NewFactory() *OptionFactory {
	return &OptionFactory{}
}

func (factory *OptionFactory) GetFactory(operatorName string) Option {
	switch operatorName {
	case "+":
		return &Add{}
	case "*":
		return &Multiple{}
	default:
		return nil
	}
}

func main() {
	operator := NewFactory().GetFactory("+")
	fmt.Println(operator.Operator(4, 5))
	operator = NewFactory().GetFactory("*")
	fmt.Println(operator.Operator(4, 5))
	operator = NewFactory().GetFactory("-")
	if operator == nil {
		fmt.Printf("NIL")
	}
}
