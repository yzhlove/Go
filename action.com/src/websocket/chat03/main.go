package main

import (
	"net/http"
	"time"
	adapter "websocket/chat03/socket_handle"

	"github.com/gorilla/websocket"
)

var HeartMessage = []byte("heartbeat")

var upgrade = websocket.Upgrader{
	//跨越允许
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//一个简单的websocket服务器
func wsHandle(w http.ResponseWriter, r *http.Request) {
	//协议转换
	var webSocket *websocket.Conn
	var err error
	var apt *adapter.WsAdapter
	var data []byte

	if webSocket, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	if apt, err = adapter.InitConn(webSocket); err != nil {
		apt.Close()
		return
	}
	go heartBeat(apt, HeartMessage)
	for {
		if data, err = apt.ReadMessage(); err != nil {
			break
		}
		if err = apt.WriteMessage(data); err != nil {
			break
		}
	}
	apt.Close()
}

//心跳包
func heartBeat(apt *adapter.WsAdapter, msg []byte) {
	for {
		if err := apt.WriteMessage(msg); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandle)
	_ = http.ListenAndServe("0.0.0.0:7777", nil)
}
