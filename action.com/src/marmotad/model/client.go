package model

import (
	"sync"
	"time"

	"github.com/davyxu/cellnet"
)

var (
	//ClientList 玩家列表
	ClientList      map[int64]*Client
	clientListMutex sync.Mutex
)

//Client 客户端信息
type Client struct {
	Session     cellnet.Session //session
	UID         string          //用户ID
	SID         string          //SID
	time        time.Timer      //定时器
	clientMutex sync.Mutex      //互斥锁
	LogicCheck  bool            //登陆状态
	CTime       time.Time       //当前时间
}

func init() {
	ClientList = make(map[int64]*Client)
}
