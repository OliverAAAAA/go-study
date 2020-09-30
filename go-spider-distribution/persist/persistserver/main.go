package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"oliver/study/go-spider-distribution/config"
	"oliver/study/go-spider-distribution/persist"
	"oliver/study/go-spider-distribution/rpcsupport"
)

var port = flag.Int("port", 0, "listen port addr")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must have port")
		return
	}
	ServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex_zhenai)
}

func ServeRpc(host, index string) error {

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  "spider_user",
	})
}
