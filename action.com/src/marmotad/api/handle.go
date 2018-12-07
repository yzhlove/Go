package api

import (
	"marmotad/proto"
	"marmotad/model"

	"github.com/davyxu/cellnet"
)

//Handle 消息处理
func Handle(ev cellnet.Event) {

	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted:
		model.NewClient(ev.Session())
	case *cellnet.SessionClosed:
		_, ok := model.ClientList[ev.Session().ID()]
		if ok {
			model.RemoveClient(ev.Session().ID())
		}
	case *proto.LOGIN:
		client , ok := model.ClientList[ev.Session().ID()]
		if !ok{
			//client在客户端列表没有找到
			return 
		}
		if client.Login(msg) {

		} else {
			
		}


}
