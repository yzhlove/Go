package main

import (
	"flag"
	"fmt"
	"journey/chat35_distributed/config"
	"journey/chat35_distributed/persist"
	"journey/chat35_distributed/rpcsupport"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "input port")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("port is %d invaild \n", *port)
		return
	}
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func ServeRpc(host, index string) error {
	var (
		client *elastic.Client
		err    error
	)

	if client, err = elastic.NewClient(elastic.SetSniff(false)); err != nil {
		return err
	}
	if err = rpcsupport.ServerRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	}); err != nil {
		return err
	}
	return nil
}
