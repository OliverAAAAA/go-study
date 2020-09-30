package zhenai

import (
	"fmt"
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/model"
	"regexp"
	"strings"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	idUrlRe   = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	imgUrl    = `https://photo.zastatic.com/images/photo/[0-9]+/%s/[0-9]+.*(\\?)(.*)`
)

func ParseCity(contents []byte, url string) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		user := model.User{}
		//城市
		//fmt.Printf("User %s \n", string(m[2]))
		name := string(m[2])
		id := extractString([]byte(m[1]), idUrlRe)
		imgUrlRe := regexp.MustCompile(fmt.Sprintf(imgUrl, id))
		imgUrl := imgUrlRe.FindAllSubmatch(contents, -1)
		user.Name = name
		user.ImgUrl = subImgUrl(string(imgUrl[0][0]))
		result.Items = []engine.Items{
			{
				Url:     string(m[1]),
				Type:    "zhenai",
				Id:      id,
				Payload: user,
			},
		}
		//url
		result.Requests = append(result.Requests, engine.MyRequest{
			Url:    string(m[1]),
			Parser: engine.NilParser{},
			//ParserFunc: func(c []byte) engine.ParseResult {
			//	//闭包
			//	return ParseProfile(c, name)
			//},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//城市
		//fmt.Printf("User %s \n", string(m[2]))
		//url
		result.Requests = append(result.Requests, engine.MyRequest{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
func subImgUrl(url string) string {
	tracer := url
	comma := strings.Index(tracer, `?`)
	return tracer[0:comma]
}
