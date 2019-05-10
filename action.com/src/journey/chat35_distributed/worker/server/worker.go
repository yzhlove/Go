package main

import (
	"flag"
	"fmt"
	"journey/chat35_distributed/rpcsupport"
	"journey/chat35_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "input port")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("port is %d invaild \n", *port)
		return
	}
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
