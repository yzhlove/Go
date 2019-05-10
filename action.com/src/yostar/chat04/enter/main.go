package main

import (
	"fmt"
	pb "yostar/chat04/proto"

	"github.com/golang/protobuf/proto"
)

func main() {

	var (
		bytes []byte
		err   error
	)

	if bytes, err = Encode(); err != nil {
		panic(err)
	}

	Decode(bytes)

}

func Encode() ([]byte, error) {
	var status pb.UserStatus
	status = pb.UserStatus_ONLINE

	userInfo := &pb.UserInfo{
		Id:     int32(10),
		Name:   "love",
		Status: status,
	}
	return proto.Marshal(userInfo)
}

func Decode(data []byte) {
	var (
		userInfo = &pb.UserInfo{}
		err      error
	)
	if err = proto.Unmarshal(data, userInfo); err != nil {
		panic(err)
	}
	fmt.Printf("Decode UserInfo:%v \n", userInfo)

}
