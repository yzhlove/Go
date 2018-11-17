package main

import "errors"

// go多值返回

func div(x,y int) (int ,error) {

	if y == 0 {
		return 0,errors.New("Div is Zero")
	}
	return x / y ,nil
}


func main() {

	value ,log := div(2,0)
	println("value = ",value," log = ",log)
}