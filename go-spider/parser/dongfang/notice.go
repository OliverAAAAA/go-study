package dongfang

import (
	"fmt"
	"github.com/antchfx/xquery/html"
	"oliver/study/go-spider/engine"
	"strings"
)

func ParseNotice(contents []byte) engine.ParseResult {
	doc, err := htmlquery.Parse(strings.NewReader(string(contents[:])))
	if err != nil {
		panic(err)
	}
	find := htmlquery.Find(doc, "//*[@class='overflow']")
	for _, node := range find {
		data := node.Data
		fmt.Printf("node data :%s", data)
	}
	result := engine.ParseResult{}
	return result
}
