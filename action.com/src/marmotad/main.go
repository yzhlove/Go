package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"

	"marmotad/api"
	"marmotad/common"

	"github.com/davyxu/golog"

	"github.com/davyxu/cellnet/proc"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
)

var host string
var port string
var mode string

var log = golog.New("游戏服务")

func init() {
	golog.SetLevelByString("gorillawsproc", "info")

	flag.StringVar(&host, "h", "0.0.0.0", "ip or hostname")
	flag.StringVar(&port, "P", "2333", "port")
	flag.StringVar(&mode, "m", "online", "debug or online")

	flag.Parse()

	check := true
	for _, m := range common.RunMode {
		if mode == m {
			check = false
		}
	}
	if check {
		log.Errorf("游戏仅支持 debug or online 模式, %s 不在支持列表!", mode)
		os.Exit(0)
	}
	log.Infof("当前游戏运行的额模式:%s \n", mode)
	common.SetMode(mode)
}

func main() {

	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", fmt.Sprintf("http://%s:%s/", host, port), queue)
	proc.BindProcessorHandler(p, "gorillaws.ltv", api.Handle)
	p.Start()
	queue.StartLoop()
	queue.Wait()
}
