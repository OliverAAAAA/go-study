package worker

import "oliver/study/go-spider/engine"

type SpiderService struct {
}

func (SpiderService) Process(req MyRequest, result *ParseResult) error {
	engineRequest, err := DeSerializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineRequest)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
