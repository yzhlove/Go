package main

import (
	"errors"
	"fmt"
	"reflect"
)

//使用反射编写万能程序

type Employee struct {
	EmployeeID uint32
	Name       string
	Age        int
}

type Custom struct {
	CustomID uint32
	Name     string
	Age      int
}

func main() {

	setting := map[string]interface{}{"Name": "love", "Age": int(18)}
	e := new(Employee)
	c := new(Custom)
	_ = full(e, setting)
	_ = full(c, setting)

	fmt.Println(e)
	fmt.Println(c)
}

func full(element interface{}, setting map[string]interface{}) error {
	et := reflect.TypeOf(element)
	if !(et.Kind() == reflect.Ptr && et.Elem().Kind() == reflect.Struct) {
		return errors.New("type err")
	}
	if setting == nil {
		return errors.New("map not nil")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range setting {
		if field, ok = reflect.ValueOf(element).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			//设置值
			reflect.ValueOf(element).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
