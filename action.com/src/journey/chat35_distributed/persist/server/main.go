package main

import (
	"journey/chat35_distributed/persist"
	"journey/chat35_distributed/rpcsupport"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal(ServeRpc(":1234", "dating_profile"))
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
