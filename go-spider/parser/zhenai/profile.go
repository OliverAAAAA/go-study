package zhenai

import (
	"oliver/study/go-spider/engine"
	"oliver/study/go-spider/model"
	"regexp"
	"strconv"
)

const re1 = `<div class="m-btn purple" [^>]*>([^<]+)</div>`
const re2 = `<div class="m-btn pink" [^>]*>([^<]+)</div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	resp := regexp.MustCompile(re1)
	info1 := resp.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	profile := model.Profile{}
	profile.Name = name
	for i := 0; i < len(info1); i++ {
		profile.Marriage = string(info1[i][1])
	}
	profile.Marriage = string(info1[0][1])
	profile.Age, _ = strconv.Atoi(string(info1[1][1]))
	profile.Xinzuo = string(info1[2][1])
	profile.Height, _ = strconv.Atoi(string(info1[3][1]))
	profile.City = string(info1[4][1])
	profile.Income = string(info1[5][1])
	profile.Occupation = string(info1[6][1])
	profile.Education = string(info1[7][1])
	//result.Items = append(result.Items, profile)
	return result
}
