package api

import (
	"marmotad/model"
	"marmotad/proto"

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
		client, ok := model.ClientList[ev.Session().ID()]
		if !ok {
			return
		}
		ack := new(proto.ACK)
		if client.Login(msg) {
			client.LogicCheck = false
			if utmp, ok := model.UserList[client.UID]; ok {
				ack.ErrorCode = 0
				ack.Result = utmp
				ev.Session().Send(ack)
				return
			}
		}
		ack.ErrorCode = 10002
		ack.ErrorReason = "登陆失败"
		ev.Session().Send(ack)
	}
}
