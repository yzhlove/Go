package main

import (
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/json"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"github.com/davyxu/cellnet/util"
	"github.com/davyxu/golog"
)

var log = golog.New("WebSocket ==>")

//TestPing 数据结构
type TestPing struct {
	Action string `json:"action"`
	Status int    `json:"status"`
}

func init() {
	//注册
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*TestPing)(nil)).Elem(),
		ID:    int(util.StringHash("main.TestPing")),
	})
}

func main() {
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", "http://0.0.0.0:8801/", queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			// log.Debugf("Accept: %+v \n", ev.Message())
		case *cellnet.SessionClosed:
			// log.Debugf("Closed: %v \n", ev.Session().ID())
		case *TestPing:
			log.Debugf(" ===> %+v \n", msg)
			ev.Session().Send(&TestPing{
				Action: msg.Action,
				Status: msg.Status + 1,
			})
		}
	})
	p.Start()
	queue.StartLoop()
	queue.Wait()
}
