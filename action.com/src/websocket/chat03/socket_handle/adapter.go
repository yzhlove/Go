package socket_handle

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	closeConnErr = errors.New("connection is close")
)

type WsAdapter struct {
	ws      *websocket.Conn
	read    chan []byte
	write   chan []byte
	notify  chan byte
	mutex   sync.Mutex
	isClose bool
}

func InitConn(ws *websocket.Conn) (*WsAdapter, error) {
	adapter := &WsAdapter{
		ws:     ws,
		read:   make(chan []byte, 1024),
		write:  make(chan []byte, 1024),
		notify: make(chan byte, 1),
	}
	go adapter.readLoop()
	go adapter.writeLoop()
	return adapter, nil
}

func (apt *WsAdapter) ReadMessage() (data []byte, err error) {
	select {
	case data = <-apt.read:
	case <-apt.notify:
		err = closeConnErr
	}
	return
}

func (apt *WsAdapter) WriteMessage(data []byte) (err error) {
	select {
	case apt.write <- data:
	case <-apt.notify:
		err = closeConnErr
	}
	return
}

func (apt *WsAdapter) Close() {
	_ = apt.ws.Close()
	apt.mutex.Lock()
	if !apt.isClose {
		close(apt.notify)
		apt.isClose = true
	}
	apt.mutex.Unlock()
}

func (apt *WsAdapter) readLoop() {
	var data []byte
	var err error
exitLoop:
	for {
		if _, data, err = apt.ws.ReadMessage(); err != nil {
			break exitLoop
		}
		select {
		case apt.read <- data:
		case <-apt.notify:
			break exitLoop
		}
	}
	apt.Close()
}

func (apt *WsAdapter) writeLoop() {
	var data []byte
	var err error
exitLoop:
	for {
		select {
		case data = <-apt.write:
		case <-apt.notify:
			break exitLoop
		}
		if err = apt.ws.WriteMessage(websocket.TextMessage, data); err != nil {
			break exitLoop
		}
	}
	apt.Close()
}
