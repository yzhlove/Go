package main

import (
	"fmt"
	"net/url"
	"strings"
)

//Go URL 参数解析

func main() {

	urlString := "https://localhost:8080@root:443/login/userLogin?name=yzh&name=xjj&age=16&birthday=19961224#once"

	u , err := url.Parse(urlString)
	if err == nil {
		fmt.Println(u)
	}

	ump , err := url.ParseRequestURI(urlString)
	if err == nil {
		fmt.Println(ump)
	}

	//得到Scheme
	fmt.Println(u.Scheme)
	fmt.Println(u.String())

	//得到认证信息
	user := u.User
	fmt.Println(user)

	userName := user.Username()
	userPasd,_ := user.Password()
	userString:=user.String()

	fmt.Printf("userName:%v userPasswd:%v userString:%v \n",userName,userPasd,userString)

	fmt.Println()

	//Host端口信息
	host := u.Host
	fmt.Println("Host = " + host)
	hostArr := strings.Split(host,":")
	fmt.Printf("hostArr:%v \n",hostArr)

	//Path信息
	path := u.Path
	fmt.Println("path = " + path)

	fmt.Println()

	//解析uri
	urlMapString := u.RawQuery
	fmt.Printf("urlMap:%T %v \n",urlMapString,urlMapString)

	urlMap , err := url.ParseQuery(urlMapString)
	if err == nil {
		for key , value := range urlMap {
			fmt.Printf("params: key :%v value :%v \n",key,value)
		}
	}

	fmt.Println()

	urlMapQuery := u.Query()
	fmt.Printf("urlMapQuery:%T %v \n",urlMapQuery,urlMapQuery)
	for key , value := range urlMapQuery {
		fmt.Printf("query key :%v value:%T \t %v \n",key,value,value)
	}

	fmt.Println()

	//添加value
	tempValue := url.Values{}
	tempValue.Set("name","xjj")
	tempValue.Set("name","yzh")
	tempValue.Set("name","xyj")
	tempValue.Set("age","18")
	tempValue.Set("birthday","19961224")
	tempValue.Add("birthday","19960523")

	fmt.Printf("urlEncode =  %v \n",tempValue.Encode())

	tempValue.Del("name")

	fmt.Printf("urlEncode =  %v \n",tempValue.Encode())

	fmt.Printf("birthdayValue: %v \n",tempValue.Get("birthday"))



}
