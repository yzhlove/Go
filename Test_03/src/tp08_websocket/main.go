package main

import (
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/cellnet/util"

	"fmt"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"

	"github.com/davyxu/golog"
)

//WebSocket服务器

var (
	log = golog.New("websocket_server")
)

//TestEchoACK 这是一个注释
type TestEchoACK struct {
	Msg   string `json:"msg"`
	Value int32  `json:"value"`
}

//ResultEchoACK 回应类型
type ResultEchoACK struct {
	ErrorCode int         `json:"error_code"`
	Action    string      `json:"action"`
	Result    interface{} `json:"result"`
}

//String 对外方法
func (test *TestEchoACK) String() string {
	return fmt.Sprintf("%+v \n", *test)
}

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("main.TestEchoACK")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*ResultEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("main.ResultEchoACK")),
	})
}

func main() {

	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", "http://0.0.0.0:8801/", queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugf("Accept: %+v \n", ev.Message())
		case *cellnet.SessionClosed:
			log.Debugf("Closed: %v \n", ev.Session().ID())
		case *TestEchoACK:
			log.Debugf(" ===> %+v \n", msg)

			data := make(map[string]string)
			data["see"] = "what are you doing"
			data["value"] = "12345"

			ev.Session().Send(&ResultEchoACK{
				ErrorCode: 3,
				Action:    "microphone",
				Result:    data,
			})

		}
	})
	p.Start()
	queue.StartLoop()
	queue.Wait()
}
