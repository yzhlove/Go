package model

/*
客户单列表操作
	创建客户端
	删除客户端
*/

import (
	"time"

	"github.com/davyxu/cellnet"
)

//NewClient 创建新的客户端
func NewClient(session cellnet.Session) {

	client := &Client{
		Session:    session,
		CTime:      time.Now(),
		LogicCheck: true,
	}
	//将客户端加入客户端列表
	clientListMutex.Lock()
	ClientList[session.ID()] = client
	clientListMutex.Unlock()
}

//RemoveClient 从客户单列表移除客户端
func RemoveClient(id int64) {
	if client, ok := ClientList[id]; ok {
		//从玩家列表清楚

		//从客户端列表清楚
		clientListMutex.Lock()
		delete(ClientList, id)
		clientListMutex.Unlock()

		//关闭连接
		client.Session.Close()
	}
}
