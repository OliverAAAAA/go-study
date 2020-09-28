package zhenai

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {

	content, _ := ioutil.ReadFile("city_test_data.html")
	//fmt.Println(string(content))
	result := ParseCity(content,"")
	//var imgUrl = `https://photo.zastatic.com/images/photo/[0-9]+/%s/[0-9]+.*(\\?)(.*)`
	//var imgUrl2 = `https://photo.zastatic.com/images/photo/[0-9]+/1682132368/[0-9]+.*(\\?).*`
	//re := fmt.Sprintf(imgUrl, "1682132368")
	//fmt.Println(re)
	//imgUrlRe := regexp.MustCompile(re)
	//find := imgUrlRe.FindAllSubmatch(content, -1)
	//fmt.Println(string(find[0][0]))
	//for _, u := range find {
	//
	//	fmt.Println(subImgUrl(string(u[0])))
	//
	//	//user.ImgUrl=string(u[1])
	//}

	fmt.Printf(string(content))
	fmt.Printf("item : %#v\n", result)
}
