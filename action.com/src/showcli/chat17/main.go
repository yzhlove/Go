package main

import (
	"fmt"
	"showcli/chat17/msg"
)

func main() {

	m := msg.Message{ID: 123, Msg: "what are you doing", Ts: 456}
	data, _ := m.MarshalMsg(nil)
	fmt.Println(string(data))

	tmp := &msg.Message{}
	distance, _ := tmp.UnmarshalMsg(data)
	fmt.Printf("tmp: %v \n distance: %v \n", tmp, distance)

}
