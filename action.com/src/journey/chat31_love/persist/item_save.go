package persist

import (
	"context"
	"fmt"
	"journey/chat31_love/engine"
	"log"

	"github.com/pkg/errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

func ItemSave() chan engine.Item {
	out := make(chan engine.Item, 128)
	go func() {
		itemCount := 0
		for {
			it := <-out
			log.Printf("Item:%d %+v\n", itemCount, it)
			itemCount++
			_, _ = save(it)
		}
	}()
	return out
}

//操作elastic search
func save(item engine.Item) error {
	var (
		client *elastic.Client
		err    error
	)
	if client, err = elastic.NewClient(elastic.SetSniff(false)); err != nil {
		panic(err)
	}

	if item.Type == "" {
		return errors.New("must not empty")
	}

	indexService := client.Index().
		Index("dating_profile").Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	if _, err = indexService.Do(context.Background()); err != nil {
		fmt.Printf("Item Err: %+v \n", item)
	}

	return nil
}
