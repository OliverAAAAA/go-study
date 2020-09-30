package main

import (
	"flag"
	"log"
	"net/rpc"
	"oliver/study/go-spider-distribution/persist/persistclient"
	"oliver/study/go-spider-distribution/rpcsupport"
	"oliver/study/go-spider-distribution/worker/workerclient"
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/parser/zhenai"
	"oliver/study/go-spider/scheduler"
	"strings"
)

var (
	itemSaverHost= flag.String("itemsaver_host","","itemsaver host")
	workerHosts = flag.String("worker_hosts","","worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	//rpc请求的保存
	itemChan, err := persistclient.ItemSaver(*itemSaverHost)
	if err != nil {
		log.Panicf("start elastic error %v", err)
	}
	pool := createWorkerClientPool(strings.Split(*workerHosts, ","))

	processor := workerclient.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.MyRequest{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(zhenai.ParseCityList, "ParseCityList"),
	})
	//e.Run(engine.MyRequest{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

func createWorkerClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
		} else {
			log.Println("create worker client pool err")
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for  {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
