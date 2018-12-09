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
		UID    string `json:"uid"`
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
	fmt.Printf("retUser:%v \n", retUser)
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

			userListMutex.Lock()
			if _, ok := UserList[u.UID]; ok {
				if cid, ok := UserClient[u.UID]; ok {
					if client, ok := ClientList[cid]; ok {
						ack := new(proto.ACK)
						ack.ErrorCode = 10003
						ack.ErrorReason = "账号异地登陆"

						client.Session.Send(ack)
						client.Session.Close()

						//删除客户端列表
						clientListMutex.Lock()
						delete(ClientList, cid)
						clientListMutex.Unlock()
					}
				}
			}
			UserClient[u.UID] = client.Session.ID()
			UserList[u.UID] = user
			client.UID = u.UID
			userListMutex.Unlock()
		} else if u.AI == 2 {
			aiUser := new(User)
			aiUser.AUID = u.UID
			aiUser.AI = u.AI
			aiUser.Avatar = u.Avatar
			aiUser.SEX = u.SEX
			aiUser.UID = u.UID
			aiUser.Name = u.Name
			aiUser.RID = fmt.Sprintf("%d", retUser.RoomID)
			aiUser.SID = u.SID
			aiUser.Ready = false
			aiUser.IsDiamond = retUser.IsDiamond
			aiUser.Addr = msg.ADDR
			aiUser.App = retUser.Code
			//将AI用户加入用户列表
			userListMutex.Lock()
			UserList[u.UID] = aiUser
			client.UID = u.UID
			userListMutex.Unlock()
		}
	}
	//查找当前用户
	if utp, ok := UserList[client.UID]; !ok || utp.UID == "" {
		return false
	}
	if retUser.RoomID == 0 {
		return false
	}
	return true
}

func getLoginAddrString(msg *proto.LOGIN) (string, bool) {

	ts := fmt.Sprintf("%d", time.Now().Unix())

	signal := map[string]string{
		"sid":  msg.SID,
		"code": msg.CODE,
		"id":   msg.GameHistoryID,
		"ts":   ts,
	}

	utemp, err := url.ParseRequestURI(common.GetLoginAddr(msg.ADDR))
	if err != nil {
		log.Errorf(err.Error())
		return "", false
	}

	urlMap, err := url.ParseQuery(utemp.RawQuery)
	if err != nil {
		log.Errorf(err.Error())
		return "", false
	}

	urlMap.Set("sid", msg.SID)
	urlMap.Set("code", msg.CODE)
	urlMap.Set("id", msg.GameHistoryID)
	urlMap.Set("ts", ts)
	urlMap.Set("h", util.Sign(signal, common.SECRET))

	// for k, v := range signal {
	// 	urlMap.Set(k, v)
	// }
	utemp.RawQuery = urlMap.Encode()
	return utemp.String(), true
}

func getRemoteUser(loginURL string) (*remoteUser, bool) {

	fmt.Printf("URL:%v \n", loginURL)

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
	res.Body.Close()
	//解析返回的参数
	remoteUser := new(remoteUser)
	err = json.Unmarshal(result, remoteUser)
	if err != nil {
		return nil, false
	}
	return remoteUser, true
}
