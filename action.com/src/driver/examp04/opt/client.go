package opt

import (
	"driver/examp04/ssdb"
)

//NewClient 新建一个ssdb客户端
func NewClient() *ssdb.Client {

	ip := "127.0.0.1"
	port := 8888

	db, err := ssdb.Connect(ip, port)
	if err != nil {
		panic(err)
	}
	return db
}
