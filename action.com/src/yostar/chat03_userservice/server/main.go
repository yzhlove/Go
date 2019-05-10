package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)
import pb "yostar/chat03_userservice/proto"

type UserInfo struct {
	ID    int
	Name  string
	Age   int
	Title []string
}

const host = ":1234"

func findUserByName(name string) (*UserInfo, error) {
	if name == "love" {
		return &UserInfo{
			ID:    888,
			Name:  "Love yzh",
			Age:   22,
			Title: []string{"聪明帅气", "活波可爱", "湖北省十佳杰出青年"},
		}, nil
	}
	return nil, errors.New("not find")
}

func (info *UserInfo) GetUserInfo(ctx context.Context, req *pb.UserReq) (*pb.UserResp, error) {
	var (
		user *UserInfo
		err  error
	)
	if user, err = findUserByName(req.Name); err != nil {
		return nil, err
	}
	return &pb.UserResp{
		Id:    int32(user.ID),
		Name:  user.Name,
		Age:   int32(user.Age),
		Title: user.Title,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[server start: %s]\n", host)
	server := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(server, &UserInfo{})
	_ = server.Serve(listen)
}

/*
1.创建TCP端口监听
2.实例化grpc服务
3.在grpc上注册服务
4.启动grpc服务
*/
