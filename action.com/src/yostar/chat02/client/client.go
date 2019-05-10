package main

import (
	"context"
	"log"
	"yostar/chat02/protc"

	"google.golang.org/grpc"
)

const address = "localhost:9988"

func main() {
	var (
		conn   *grpc.ClientConn
		err    error
		client protc.ToUpperClient
		resp   *protc.UpperResp
	)
	if conn, err = grpc.Dial(address, grpc.WithInsecure()); err != nil {
		panic(err)
	}
	defer conn.Close()
	client = protc.NewToUpperClient(conn)
	sendMsg := "hello world"

	if resp, err = client.Upper(context.Background(), &protc.UpperReq{Name: sendMsg}); err != nil {
		log.Fatalf("could not greet: %v ", err)
	}
	log.Printf("Response: %s ", resp.Message)

}
