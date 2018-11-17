package main

import (
	"fmt"
	"runtime"
)

//Goexit

func main() {

	exit := make(chan struct{})
	go func() {
		defer close(exit)
		defer fmt.Println("a exit ...")

		func() {
			defer func() {
				fmt.Println("b", recover() == nil)
			}()

			func() {
				fmt.Println("c ...")
				runtime.Goexit()
				fmt.Println("c deone ...")
			}()

			fmt.Println(" b done ...")

		}()

		fmt.Println("a done ... ")

	}()

	<-exit

	fmt.Println("main.exit ...")

}
