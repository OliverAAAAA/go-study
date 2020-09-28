package engine

/**
请求对象
*/
type MyRequest struct {
	Url        string
	ParserFunc func([]byte,string) ParseResult
}

/**
解析器结果
*/
type ParseResult struct {
	Requests []MyRequest
	Items    []Items
}

type Items struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

func NilParser([]byte,string) ParseResult {
	return ParseResult{}
}

