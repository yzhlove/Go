package main

import "fmt"

//接口类型转换

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Brid struct{}

type Pig struct{}

func (bird Brid) Fly() {
	fmt.Println("bird fly.")
}

func (bird Brid) Walk() {
	fmt.Println("bird walk.")
}

func (pig Pig) Fly() {
	fmt.Println("pig fly.")
}

func main() {

	animals := make(map[string]interface{})
	animals["bird"] = new(Brid)
	animals["pig"] = new(Pig)

	for name, inter := range animals {
		fmt.Printf("name:%s inter:%+v \n", name, inter)
		if fly, isFly := inter.(Flyer); isFly {
			fly.Fly()
		}
		if walk, isWalk := inter.(Walker); isWalk {
			walk.Walk()
		}
	}

}
