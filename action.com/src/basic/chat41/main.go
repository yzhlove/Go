package main

import (
	"basic/chat41/event"
	"fmt"
)

//一个简单的事件调用

type Actor struct{}

func (actor *Actor) OnEvent(param interface{}) {
	fmt.Println("onEvent:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("GlobalEvent:", param)
}

func main() {

	actor := new(Actor)

	event.RegisterEvent("actor", actor.OnEvent)
	event.RegisterEvent("actor", GlobalEvent)
	event.CallEvent("actor", map[string]string{"hello": "World"})

}
