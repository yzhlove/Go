package main

import (
	"fmt"
	"journey/chat35_distributed/config"
	"journey/chat35_distributed/rpcsupport"
	"journey/chat35_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
