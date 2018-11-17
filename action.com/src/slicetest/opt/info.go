package opt

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

//UserInfo 用户信息
type UserInfo struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	URI      string `json:"url"`
	Birthday string `json:"birthday"`
}

const filePath = "/Users/love/WorkSpace/Go/action.com/src/slicetest/data/data.json"
const maxSize = 100

//InitUserData 初始化用户数据
func InitUserData() ([]*UserInfo,error) {

	file , err := os.Open(filePath)
	if err != nil {
		log.Fatalln("OpenFileErr:",err)
		return nil , err
	}
	//无论如何关闭文件流
	defer file.Close()
	userIfs := make([]*UserInfo,0,maxSize)
	err = json.NewDecoder(file).Decode(&userIfs)
	return userIfs,err
}


//ShowString 用户信息显示
func (userInfo *UserInfo) ShowString() string {

	var str string

	str += "+---------------------------------+\n"
	str += "| Name = " + userInfo.Name + "\n"
	str += "| Age = " + strconv.Itoa(userInfo.Age) + "\n"
	str += "| URI = " + userInfo.URI + "\n"
	str += "| Birthday = " + userInfo.Birthday + "\n"
	str += "+---------------------------------+\n"

	return str
}
