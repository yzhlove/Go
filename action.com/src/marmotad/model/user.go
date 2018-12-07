package model

import (
	"sync"
	"time"
)

//User 用户信息
type User struct {
	AI        int    `json:"ai"`
	Avatar    string `json:"avatar"`
	SEX       int    `json:"sex"`
	UID       int    `json:"UID"`
	Name      string `json:"user_name"`
	AUID      int    `json:"-"`
	RID       string `json:"-"`
	Ready     bool   `json:"ready"`
	SID       string `json:"-"`
	EnterFrom string `json:"-"`
	SiteTime  time.Time
	mutex     sync.Mutex
	Addr      string
	IsDiamond int
	App       string `json:"-"`
}

var (
	//UserList 玩家列表
	UserList      map[int]*User
	userListMutex sync.Mutex
	//UserClient User与Client对应
	UserClient map[int]int64
)

func init() {
	UserList = make(map[int]*User)
	UserClient = make(map[int]int64)
}
