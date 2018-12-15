package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//荣庆物流接口请求

//Request 请求头
type Request struct {
	WayCode   string
	ShipperID int
	WayTime   string
}

/*
WayCode: 'CO2018121108408',
ShipperID:28955,
WayTime:'2018-12-14T12:18:22.590Z'
*/

const url = "http://vip.rokin.cn:8088/api/TMS/GetRokinGps"

func main() {

	request := &Request{
		WayCode:   "CO2018121108408",
		ShipperID: 28955,
		WayTime:   "2018-12-14T12:18:22.590Z",
	}

	requestByteArr, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Json Err ")
		return
	}
	//application/x-www-form-urlencoded | application/json;charset=utf-8
	response, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(requestByteArr))
	if err != nil {
		fmt.Println("response Err ")
		return
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("io read Err ")
		return
	}
	response.Body.Close()
	fmt.Printf("response : %v \n", string(content))
}
