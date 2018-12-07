package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marmotad/common"
	"marmotad/proto"
	"marmotad/util"
	"net/http"
	"net/url"
	"time"
)

/*
客户端登陆
	Client -> User
*/

//AppServer Login返回的数据
type remoteUser struct {
	ErrorCode   int    `json:"error_code"`
	ErrorReason string `json:"error_reason"`
	ErrorURL    string `json:"error_url"`
	RoomID      int    `json:"game_room_uid"`
	Code        string `json:"code"`
	IsDiamond   int    `json:"scene_type"`
	Users       []struct {
		ID     int    `json:"id"`
		UID    int    `json:"uid"`
		SID    string `json:"sid"`
		SEX    int    `json:"sex"`
		Avatar string `json:"avatar_small_url"`
		Name   string `json:"nickname"`
		AI     int    `json:"user_type"`
	} `json:"users"`
}

//Login 客户端登陆
func (client *Client) Login(msg *proto.LOGIN) bool {

	loginURL, ok := getLoginAddrString(msg)
	if !ok {
		return false
	}
	retUser, ok := getRemoteUser(loginURL)
	if !ok {
		return false
	}
	if retUser.ErrorCode != 0 {
		log.Errorln("error_reason:" + retUser.ErrorReason)
		return false
	}
	//初始化用户数据
	user := new(User)
	user.Ready = false
	user.RID = fmt.Sprintf("%d", retUser.RoomID)
	user.EnterFrom = msg.EnterFrom
	user.Addr = msg.ADDR
	user.IsDiamond = retUser.IsDiamond
	user.App = retUser.Code

	for _, u := range retUser.Users {
		if u.SID == msg.SID {
			user.AI = u.AI
			user.Avatar = u.Avatar
			user.SEX = u.SEX
			user.UID = u.UID
			user.Name = u.Name
		} else if u.AI == 2 {

		}
	}

	return true
}

func getLoginAddrString(msg *proto.LOGIN) (string, bool) {

	ts := fmt.Sprintf("%d", time.Now().Unix())

	signal := make(map[string]string)
	signal["sid"] = msg.SID
	signal["code"] = msg.CODE
	signal["id"] = msg.GameHistoryID
	signal["ts"] = ts

	utemp, err := url.ParseRequestURI(common.GetLoginAddr(msg.ADDR))
	if err == nil {
		log.Errorf(err.Error())
		return "", false
	}

	urlMap, err := url.ParseQuery(utemp.RawQuery)
	if err == nil {
		log.Errorf(err.Error())
		return "", false
	}

	signalStr := util.Sign(signal, common.SECRET)
	signal["h"] = signalStr

	for k, v := range signal {
		urlMap.Set(k, v)
	}
	utemp.RawQuery = urlMap.Encode()
	return utemp.String(), true
}

func getRemoteUser(loginURL string) (*remoteUser, bool) {

	res, err := http.Get(loginURL)
	if err != nil {
		log.Errorf(err.Error())
		return nil, false
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf(err.Error())
		return nil, false
	}
	defer res.Body.Close()
	//解析返回的参数
	remoteUser := new(remoteUser)
	err = json.Unmarshal(result, remoteUser)
	if err != nil {
		return nil, false
	}
	return remoteUser, true
}
