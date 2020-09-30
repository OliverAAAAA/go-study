package worker

import (
	"errors"
	"log"
	"oliver/study/go-spider-distribution/config"
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/parser/zhenai"
)

type SerializedParser struct {
	FunctionName string
	Args         interface{}
}

//{"ParseCityList",nil}   ,



type MyRequest struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Items
	Requests []MyRequest
}

//序列化request
func SerializeRequest(r engine.MyRequest) MyRequest {
	name, args := r.Parser.Serialize()
	return MyRequest{
		Url: r.Url,
		Parser: SerializedParser{
			FunctionName: name,
			Args:         args,
		},
	}
}

//序列化result
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

//反序列化
func DeSerializeRequest(r MyRequest) (engine.MyRequest, error) {
	parser, err := deSerializeParser(r.Parser)
	if err != nil {
		return engine.MyRequest{}, err
	}
	return engine.MyRequest{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

//反序列化
func DeSerializeResult(r ParseResult) (engine.ParseResult, error) {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request, err := DeSerializeRequest(req)
		if err != nil {
			log.Printf("error deSerialize req error  %v",err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result, nil
}

func deSerializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FunctionName {
	case config.ParseCity:
		return engine.NewFuncParser(zhenai.ParseCity, config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(zhenai.ParseCityList, config.ParseCityList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("deSerializeParser error")
	}
}

