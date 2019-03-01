package client

import (
	"journey/chat31_love/engine"
	"journey/chat35_distributed/rpcsupport"
	"log"
	"net/rpc"
)

func ItemSaver(host string) (chan engine.Item, error) {
	var (
		client *rpc.Client
		err    error
	)
	if client, err = rpcsupport.NewClient(host); err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		var (
			itemCount = 0
			result    = ""
		)
		for {
			item := <-out
			log.Printf("Item Saver #%d , %v", itemCount, item)
			if err = client.Call("ItemSaveService.Save", item, &result); err != nil {
				log.Printf("Item Saver Err %s %v", result, err)
				continue
			}
			itemCount++
		}
	}()
	return out, nil
}
