package main

import (
	"oliver/study/go-spider-distribution/rpcsupport"
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/model"
	"testing"
	"time"
)

const host = ":1234"

func TestItemSaver(t *testing.T) {
	//start persistserver
	go ServeRpc(host, "spider_user")
	time.Sleep(time.Second)
	//start persistclient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	user := model.User{}
	user.ImgUrl=""
	user.Name="OliverAAAAA"
	item := engine.Items{
		Url:     "123",
		Id:      "0",
		Type:    "zhenai",
		Payload: user,
	}

	result := ""
	client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result : %s ,err : %v", result,err)
	}
}
