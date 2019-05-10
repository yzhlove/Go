package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strings"
	"yostar/chat02/protc"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
	port = ":9988"
)

type toUpperSev struct{}

func (s *toUpperSev) Upper(ctx context.Context, req *protc.UpperReq) (*protc.UpperResp, error) {
	log.Printf("[Recv:%v]\n", req)
	if req.Name == "" {
		return nil, errors.New("not nil")
	}
	sendMsg := &protc.UpperResp{Message: strings.ToUpper(req.Name)}
	return sendMsg, nil
}

func main() {
	var (
		listen net.Listener
		err    error
		server *grpc.Server
	)
	if listen, err = net.Listen("tcp", port); err != nil {
		log.Fatalln(err.Error())
	}
	server = grpc.NewServer()
	protc.RegisterToUpperServer(server, &toUpperSev{})
	reflection.Register(server)
	if err = server.Serve(listen); err != nil {
		log.Fatalln(err.Error())
	}
}
