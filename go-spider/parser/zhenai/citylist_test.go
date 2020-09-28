package zhenai

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	content, _ := ioutil.ReadFile("citylist_test_data.html")
	result := ParseCityList(content,"")

	//verify result
	const resultSize = 470
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d request;but had %d",resultSize,len(result.Requests))
	}
}
