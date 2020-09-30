package engine

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (funcName string, args interface{})
}

type ParserFunc func(contents []byte, url string) ParseResult

/**
请求对象
*/
type MyRequest struct {
	Url    string
	Parser Parser
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

type NilParser struct {
}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (funcName string, args interface{}) {
	return "NilParser", nil
}

//func NilParser([]byte, string) ParseResult {
//	return ParseResult{}
//}

type FuncParser struct {
	parser ParserFunc
	Name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f FuncParser) Serialize() (funcName string, args interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		Name:   name,
	}
}
