package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

//一个简单的websocket服务器

var (
	upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	//握手失败，关闭连接
	if conn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	_ = http.ListenAndServe("0.0.0.0:7777", nil)
}
