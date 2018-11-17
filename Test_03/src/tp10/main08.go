package main

import (
	"fmt"
	"reflect"
)

//反射

//S 集合
type Ss struct {
}

//T 集合
type Ts struct {
	Ss
}

func (Ss) Sval()  {}
func (*Ss) Sptr() {}
func (Ts) Tval()  {}
func (*Ts) Tptr() {}

func methodSet(a interface{}) {

	t := reflect.TypeOf(a)
	fmt.Printf("%+v %+v \n", t, t.NumMethod())
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}

func main() {

	var t Ts

	methodSet(t)
	fmt.Println("--------------------")
	methodSet(&t)

}

/*
==> result

main.Ts 2
Sval func(main.Ts)
Tval func(main.Ts)
--------------------
*main.Ts 4
Sptr func(*main.Ts)
Sval func(*main.Ts)
Tptr func(*main.Ts)
Tval func(*main.Ts)
*/
