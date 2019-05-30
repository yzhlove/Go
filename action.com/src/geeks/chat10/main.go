package main

import (
	"sort"

	"fmt"
)

type User struct {
	Age  uint32
	Name string
}

type UserInfoList []*User

type Users struct {
	data UserInfoList
}

func (u UserInfoList) Len() int {
	return len(u)
}

func (u UserInfoList) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserInfoList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u *Users) SearchUser(age uint32) int {

	//找到返回下标，找不到返回length
	index := sort.Search(len(u.data), func(i int) bool {
		return u.data[i].Age >= age
	})
	return index
}

func main() {
	tempMap := make(map[uint32]string, 5)
	tempMap[16] = "xjj"
	tempMap[17] = "lcm"
	tempMap[15] = "xyj"
	tempMap[18] = "fyb"
	tempMap[20] = "hxy"

	users := &Users{}

	for age, name := range tempMap {
		users.data = append(users.data, &User{Age: age, Name: name})
	}

	//排序
	sort.Sort(users.data)

	for i := 0; i < len(users.data); i++ {
		fmt.Println(users.data[i])
	}

	fmt.Println("=================")

	//查找
	fmt.Println(users.SearchUser(24))
}
