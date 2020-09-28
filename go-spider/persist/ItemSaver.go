package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
	"oliver/study/go-spider/engine"
)

type SaveItem struct {
	Name string
}

func ItemSaver(index string) (chan engine.Items, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Items)
	go func() {
		//save
		itemCount := 0
		for {
			item := <-out
			log.Printf("Got item : %v ,count %d \n", item, itemCount)
			itemCount++

			err := save(item, client,index)
			if err != nil {
				log.Printf("save error item %v err %v\n", item, err)
			}

		}
	}()
	return out,nil
}

func save(item engine.Items, client *elastic.Client,index string) (err error) {

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("must need type")
	}

	itemService := client.Index().
		Index(index).
		//Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		itemService.Id(item.Id)
	}
	_, err = itemService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
