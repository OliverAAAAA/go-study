package persistclient

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
	"oliver/study/go-spider-distribution/config"
	"oliver/study/go-spider-distribution/rpcsupport"
	"oliver/study/go-spider/engine"
)

func ItemSaver(host string) (chan engine.Items, error) {
	client, err := rpcsupport.NewClient(host)
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

			result := ""
			//Call rpc save
			client.Call(config.ItemSaverRpc, item, &result)
			//err := Save(item, persistclient,index)

			if err != nil || result != "ok" {
				log.Printf("result : %s ,err : %v", result, err)
			}

		}
	}()
	return out, nil
}

func Save(item engine.Items, client *elastic.Client, index string) (err error) {

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
