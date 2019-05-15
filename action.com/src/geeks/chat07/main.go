package main

import (
	"fmt"
	"time"
)

//时间转换测试

func main() {

	ts, err := time.Parse("2006-01-02 15:04:05", "2019-05-15 14:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println("ts => ", ts.Unix())

	tsStr := ts.Format("2006-01-02 15:04:05")
	fmt.Println("str => ", tsStr)
}
