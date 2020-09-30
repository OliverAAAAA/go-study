package persist

import (
	"github.com/olivere/elastic/v7"
	"log"
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Items, result *string) error {

	err := persist.Save(item, s.Client, s.Index)
	log.Printf("item %v saved",item)
	if err == nil {
		*result = "ok"
	}
	return err
}
