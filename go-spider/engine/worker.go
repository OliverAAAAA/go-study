package engine

import (
	"log"
	"oliver/study/go-spider/fetcher"
)

func worker(req MyRequest) (ParseResult, error) {
	log.Printf("Fetching %s\n", req.Url)
	//请求req的url
	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		log.Printf("Fetecher error fetching url :%s err:%v", req.Url, err)
		return ParseResult{}, err
	}
	//解析req的html
	return req.ParserFunc(body, req.Url), nil
}
