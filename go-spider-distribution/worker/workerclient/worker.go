package workerclient

import (
	"net/rpc"
	"oliver/study/go-spider-distribution/config"
	"oliver/study/go-spider-distribution/worker"
	"oliver/study/go-spider/engine"
)

func CreateProcessor(clientsChan chan *rpc.Client) (engine.Processor) {

	return func(request engine.MyRequest) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c:= <- clientsChan
		err2 := c.Call(config.SpiderServiceRpc, sReq, &sResult)
		if err2 != nil {
			return engine.ParseResult{}, err2
		}
		return worker.DeSerializeResult(sResult)
	}

}
