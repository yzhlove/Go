package main

import (
	"fmt"
	"reflect"
)

//字段标签
type user struct {
	name     string `USER_NAME`
	sex      byte   `USER_SEX`
	birthday string `USER_BIRTHDAY`
}

func main() {

	u := user{
		name:     "Jerry",
		sex:      1,
		birthday: "1996-12-24",
	}
	v := reflect.ValueOf(u)
	t := v.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s:%v\n", t.Field(i).Tag, v.Field(i))
	}

}
