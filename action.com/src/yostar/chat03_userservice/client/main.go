package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	pb "yostar/chat03_userservice/proto"

	"google.golang.org/grpc"
)

const host = ":1234"

var regMust = regexp.MustCompile(`\\[0-7]{3}`)

func toChineseSimple(in string) string {
	out := regMust.ReplaceAllFunc([]byte(in), func(bytes []byte) []byte {
		i, _ := strconv.ParseInt(string(bytes[1:]), 8, 0)
		return []byte{byte(i)}
	})
	return string(out)
}

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewUserInfoServiceClient(conn)
	request := &pb.UserReq{
		Name: "love",
	}
	response, err := client.GetUserInfo(context.Background(), request)
	if err != nil {
		log.Fatalf("resp err:%v \n", err)
	}
	fmt.Printf("resp:%v \n", response)

	for _, value := range response.Title {
		title := toChineseSimple(value)
		fmt.Printf("title:%v \n", title)
	}

}

/*
1。创建与grpc服务的连接
2。实例化grpc客户端
3。调用接口
*/
