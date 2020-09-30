package main

import (
	"fmt"
	"oliver/study/go-spider-distribution/config"
	"oliver/study/go-spider-distribution/rpcsupport"
	"oliver/study/go-spider-distribution/worker"
	"testing"
	"time"
)

func TestSpiderService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.SpiderService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.MyRequest{
		Url: "http://www.zhenai.com/zhenghun/zhongxian",
		Parser: worker.SerializedParser{
			FunctionName: config.ParseCity,
			Args:         "",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.SpiderServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
