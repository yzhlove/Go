package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
	Type int `json:"type" id:"100"`
}

func main() {

	ins := cat{Name: "admin", Type: 1}

	typecat := reflect.TypeOf(ins)

	fmt.Println(typecat.Name(), typecat.Kind())

	//遍历所有的结构体成员
	for i := 0; i < typecat.NumField(); i++ {
		//获取每个成员结构题的字段类型
		fieldType := typecat.Field(i)
		fmt.Printf("name:%v tag:'%v' path:%v \n", fieldType.Name, fieldType.Tag, fieldType.PkgPath)
	}

	//通过字段名找到字段类型信息
	if catType, ok := typecat.FieldByName("Type"); ok {
		//根据key判断值是否存在
		if value, ok := catType.Tag.Lookup("json"); ok {
			fmt.Println("json = ", value)
		}
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

}
