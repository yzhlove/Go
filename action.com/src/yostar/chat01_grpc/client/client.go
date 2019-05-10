package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
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
		log.Println("print ------> ", feature)
	}
}

func printFeatures(client proto.RouteGuideClient, rect *proto.Rectangle) {
	log.Fatalf("Looking for features within %v", rect)
	ctx, ctf := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctf()
	if stream, err := client.ListFeatures(ctx, rect); err != nil {
		log.Fatalf("%V ListFeatures() = _, %v \n", client, err)
	} else {
		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v Listfeatures() , _ ,%v \n", client, err)
			}
			log.Println(feature)
		}
	}
}

func runRecordRoute(client proto.RouteGuideClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	printCount := int(r.Int31n(100)) + 2
	var points []*proto.Point
	for i := 0; i < printCount; i++ {
		points = append(points, randPoint(r))
	}
	log.Printf("Trversing %d points.", len(points))
	ctx, ctf := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctf()

	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("%v RecordRoute() , _ ,%v \n", client, err)
	}
	for _, point := range points {
		if err = stream.Send(point); err != nil {
			log.Fatalf("%v send(%v) = %v \n", stream, point, err)
		}
	}
	replay, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v CloseAndRecv() got error %v ,want %v \n", stream, err, nil)
	}
	log.Printf("Route summary: %v ", replay)
}

func runRouteChat(client proto.RouteGuideClient) {
	notes := []*proto.RouteNote{
		{Location: &proto.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &proto.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &proto.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &proto.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &proto.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &proto.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	ctx, ctf := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctf()

	stream, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("%v RouteChat() = _ , %v \n", client, err)
	}
	wait := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err != io.EOF {
				close(wait)
				return
			}
			if err != nil {
				log.Fatalf("Failure to receive a note :%v ", err)
			}
			log.Printf("Got mesage %s at point(%d , %d )", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to Send a note : %v ", err)
		}
	}
	_ = stream.CloseSend()
	<-wait
}

func randPoint(r *rand.Rand) *proto.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &proto.Point{Latitude: lat, Longitude: long}
}

func main() {
	flag.Parse()
	var (
		opts   []grpc.DialOption
		err    error
		creeds credentials.TransportCredentials
		conn   *grpc.ClientConn
	)
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		if creeds, err = credentials.NewClientTLSFromFile(*caFile, *serverHostOverride); err != nil {
			log.Fatalf("Failed to create TLS err: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creeds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if conn, err = grpc.Dial(*serverAddr, opts...); err != nil {
		log.Fatalf("fail to dail:%v ", err)
	}
	defer conn.Close()
	client := proto.NewRouteGuideClient(conn)

	printFeature(client, &proto.Point{Latitude: 409146138, Longitude: -746188906})
	printFeatures(client, &proto.Rectangle{
		Lo: &proto.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &proto.Point{Latitude: 420000000, Longitude: -730000000},
	})
	runRecordRoute(client)
	runRouteChat(client)
}
