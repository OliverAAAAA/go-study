package zhenai

import (
	"oliver/study/go-spider/engine"
	"regexp"
)

//<a href="http://www.zhenai.com/zhenghun/zhumadian" data-v-2cb5b6a2>驻马店</a>
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte,url string) engine.ParseResult {
	resp := regexp.MustCompile(cityListRe)
	cityListA := resp.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range cityListA {
		//城市
		//result.Items = append(result.Items, string(m[2]))
		//url
		result.Requests = append(result.Requests, engine.MyRequest{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity,"ParseCity"),
		})
	}

	return result
}
