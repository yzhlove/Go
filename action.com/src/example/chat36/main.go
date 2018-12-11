package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

const sercet = "ca414f5145129530e061faaeb55c9dc8"

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

//测试验权算法是否正确

func sign(m map[string]string, secret string) string {

	var params []string
	for _, v := range []string{"_url", "h", "file"} {
		delete(m, v)
	}
	for k, v := range m {
		params = append(params, k+"="+v)
	}
	sort.Strings(params)
	str := strings.Join(params, "&")
	return entryMD5(entryMD5(str) + secret)

}

func entryMD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	s := ctx.Sum(nil)
	return hex.EncodeToString(s)
}

func main() {

	hashMap := make(map[string]string)
	hashMap["sid"] = "114s8acdfc1ab3f163c0a4c7fec865bd30aee5"
	hashMap["code"] = "yuewan"
	hashMap["id"] = "17564"
	hashMap["request_root"] = "http://test.momoyuedu.cn"
	// hashMap["ts"] = fmt.Sprintf("%d", time.Now().Unix())
	result := sign(hashMap, sercet)
	hashMap["h"] = result
	fmt.Printf("sercet:%v \n", result)
	fmt.Printf("hashMap:%v \n", hashMap)

	paramsArr := make([]string, 0, len(hashMap))
	for k, v := range hashMap {
		paramsArr = append(paramsArr, k+"="+v)
	}
	paramString := strings.Join(paramsArr, "&")
	url := "http://test.momoyuedu.cn/game/game_histories/detail"
	usrString := url + "?" + paramString

	response, err := http.Get(usrString)
	if err != nil {
		panic("Request ERR!")
	}
	responseText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("responseText ERR!")
	}
	response.Body.Close()
	fmt.Printf("responseText:%v \n", string(responseText))

	remoteUser := new(remoteUser)
	err = json.Unmarshal(responseText, remoteUser)
	if err != nil {
		log.Fatalln(err.Error())
		panic("remoteUser ERR!")
	}

	fmt.Printf("remoteUser == %v \n", remoteUser)
}
