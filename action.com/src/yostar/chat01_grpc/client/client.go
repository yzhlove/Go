package main

import (
	"flag"
	"yostar/chat01_grpc/proto"

	"github.com/name5566/leaf/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls                = flag.Bool("tls", false, "<TLS>")
	caFile             = flag.String("ca_file", "", "<CA> File")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "<Host Addr>")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "<- ->")
)

func main() {
	flag.Parse()
	var (
		opts  []grpc.DialOption
		err   error
		creds credentials.TransportCredentials
		conn  *grpc.ClientConn
	)
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		if creds, err = credentials.NewClientTLSFromFile(*caFile, *serverHostOverride); err != nil {
			log.Fatal("Failed to create TLS err: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if conn, err = grpc.Dial(*serverAddr, opts...); err != nil {
		log.Fatal("fail to dail:%v ", err)
	}
	defer conn.Close()
	client := proto.NewRouteGuideClient(conn)

}
