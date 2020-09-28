package dongfang

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNotice(t *testing.T) {
	content, _ := ioutil.ReadFile("notice_test_data.html")
	result := ParseNotice(content)

	fmt.Printf("item : %#v\n", result)
}