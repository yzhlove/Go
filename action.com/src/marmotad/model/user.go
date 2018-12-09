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
	UID       string `json:"UID"`
	Name      string `json:"user_name"`
	AUID      string `json:"-"`
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
	UserList      map[string]*User
	userListMutex sync.Mutex
	//UserClient User与Client对应
	UserClient map[string]int64
)

func init() {
	UserList = make(map[string]*User)
	UserClient = make(map[string]int64)
}
