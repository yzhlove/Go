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

	fmt.Println("=============================")

	testTs()

}

func testTs() {

	ts := time.Now().Unix()
	fmt.Println(ts)
	str := time.Unix(ts, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str)

	tsLocal, _ := time.LoadLocation("Local")
	//t, _ := time.Parse("2006-01-02 15:04:05", str)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", str, tsLocal)
	fmt.Println(t.Unix())

}
