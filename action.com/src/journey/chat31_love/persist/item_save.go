package persist

import (
	"context"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

func ItemSave() chan interface{} {
	out := make(chan interface{}, 128)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item:%d %v\n", itemCount, item)
			itemCount++
			if _, err := save(item); err != nil {
				log.Printf("Save Err %+v : %+v \n", item, err)
			}
		}
	}()
	return out
}

//操作elastic search
func save(item interface{}) (string, error) {
	var (
		client *elastic.Client
		err    error
		resp   *elastic.IndexResponse
	)
	if client, err = elastic.NewClient(elastic.SetSniff(false)); err != nil {
		panic(err)
	}

	if resp, err = client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).Do(context.Background()); err != nil {
		panic(err)
	}

	return resp.Id, nil
}
