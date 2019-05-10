package main

import (
	"fmt"
	"reflect"
)

type Dummy struct {
	a int
	b string
	float32
	bool
	next *Dummy
}

func main() {

	d := reflect.ValueOf(Dummy{next: &Dummy{}})
	fmt.Println("NumField:", d.NumField())
	//根据下表查找字段类型
	floatField := d.Field(2)
	fmt.Println("FieldByName = ", floatField.Type())
	//根据变量名查找字段类型
	fmt.Println("FieldByName = ", d.FieldByName("b").Type())
	//[]int{4,0}中的4表示，在dummy结构中索引值为4的成员，也就是next。next的类型为dummy，也是一个结构体，因此使用[]int{4,0}中的0继续在next值的基础上索引，结构为dummy中索引值为0的a字段，类型为int。
	fmt.Println("FieldByIndex = ", d.FieldByIndex([]int{4, 0}).Type())

}
