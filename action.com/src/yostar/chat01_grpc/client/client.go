package main

import (
	"context"
	"flag"
	"log"
	"time"
	"yostar/chat01_grpc/proto"

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

func printFeature(client proto.RouteGuideClient, point *proto.Point) {
	log.Printf("Getting feature for point (%d,%d)", point.Latitude, point.Longitude)
	ctx, ctf := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctf()
	if feature, err := client.GetFeature(ctx, point); err != nil {
		log.Fatalf("%v . GetFeatures() = , _ , %v \n", client, err)
	} else {
		log.Println(feature)
	}
}

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
			log.Fatalf("Failed to create TLS err: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if conn, err = grpc.Dial(*serverAddr, opts...); err != nil {
		log.Fatalf("fail to dail:%v ", err)
	}
	defer conn.Close()
	client := proto.NewRouteGuideClient(conn)
	printFeature(client, &proto.Point{Latitude: 409146138, Longitude: -746188906})

}
