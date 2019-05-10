package persist

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/persist"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("[-----] Item %v saved.", item)
	if err != nil {
		log.Printf("[-----]Error saving item: %v err:%v ", item, err)
	} else {
		*result = "ok"
	}
	return err
}
