package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...MyRequest) {

	var myRequest []MyRequest
	for _, r := range seeds {
		myRequest = append(myRequest, r)
	}

	for len(myRequest) > 0 {
		//使用第一个req
		req := myRequest[0]
		myRequest = myRequest[1:]
		parseResult, err := worker(req)
		if err != nil {
			continue
		}
		//将解析结果加入到队列中
		myRequest = append(myRequest, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
