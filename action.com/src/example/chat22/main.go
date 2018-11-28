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
	"time"
)

type resT struct {
	ErrorCode 	int `json:"error_code"`
	ErrorReason string `json:"error_reason"`
	ErrorURL 	string `json:"-"`
	NowAt		int64 `json:"-"`
	Result 		[]struct{
		ID int `json:"id"`
		OriginalImage string `json:"original_image"`
		DifferentImage string `json:"different_image"`
		Answer []struct{
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"answer"`
	} `json:"data"`
}

const (
	SECRET     = "ca414f5145129530e061faaeb55c9dc8"
	requestURL = "http://jump.com/game/quickspots/index"
)

func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	s := ctx.Sum(nil)
	return hex.EncodeToString(s)
}

func sign(m map[string]string, s string) string {
	params := make([]string, 0, len(m))
	for _, v := range []string{"_url", "h", "file"} {
		delete(m, v)
	}
	for k, v := range m {
		params = append(params, k+"="+v)
	}
	sort.Strings(params)
	str := strings.Join(params, "&")

	return MD5(MD5(str) + s)
}

func requestData() (string, error) {

	body := make(map[string]string)
	body["ts"] = fmt.Sprintf("%d",time.Now().Unix())
	body["number"] = "2"
	body["h"] = sign(body, SECRET)

	request := make([]string, 0, len(body))
	for k, v := range body {
		request = append(request, k+"="+v)
	}
	response, err := http.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(strings.Join(request,"&")))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {

	str, err := requestData()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//fmt.Println("response: ", str)

	rs := new(resT)
	err = json.Unmarshal([]byte(str),rs)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	fmt.Printf("resT => %+v \n",rs)
}
