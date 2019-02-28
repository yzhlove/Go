package persist

import (
	"journey/chat31_love/engine"
	"journey/chat31_love/persist"

	elastic "gopkg.in/olivere/elastic.v5"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	*result = "ok"
	return err
}
