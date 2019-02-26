package persist

import (
	"context"
	"fmt"
	"journey/chat31_love/engine"
	"log"

	"github.com/pkg/errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

func ItemSave(index string) (chan engine.Item, error) {
	var (
		client *elastic.Client
		err    error
	)
	if client, err = elastic.NewClient(elastic.SetSniff(false)); err != nil {
		return nil, err
	}

	out := make(chan engine.Item, 128)
	go func() {
		itemCount := 0
		for {
			it := <-out
			log.Printf("Item:%d %+v\n", itemCount, it)
			itemCount++
			_ = save(client, index, it)
		}
	}()
	return out, nil
}

//操作elastic search
func save(client *elastic.Client, index string, item engine.Item) error {
	var err error
	if item.Type == "" {
		return errors.New("must not empty")
	}

	indexService := client.Index().
		Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	if _, err = indexService.Do(context.Background()); err != nil {
		fmt.Printf("Item Err: %+v \n", item)
	}

	return nil
}
