package main

import (
	"log"
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/parser/zhenai"
	"oliver/study/go-spider/persist"
	"oliver/study/go-spider/scheduler"
)

func main() {

	itemChan, err := persist.ItemSaver("spider_user")
	if err != nil {
		log.Panicf("start elastic error %v", err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
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
