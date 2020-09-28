package zhenai

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	content, _ := ioutil.ReadFile("profile_test_data.html")
	result := ParseProfile(content,"")

	fmt.Printf("item : %#v\n", result.Items)
}
