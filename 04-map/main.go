package main

import (
	"fmt"
	"strings"
)

func main() {
	//元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8)
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["test1"] = 1
	fmt.Println(mapSlice)

	//值为切片的map
	var mapValue = make(map[string][]int, 8)
	v, ok := mapValue["111"]
	if ok {
		fmt.Println(mapValue["111"], v)
	} else {
		//value切片初始化
		mapValue["111"] = make([]int, 8)
		mapValue["111"][0] = 100
		mapValue["111"][1] = 200
		mapValue["111"][2] = 300
		mapValue["111"][3] = 400
	}

	for k, v := range mapValue {
		fmt.Println(k, v)
	}

	var s = "how do you do"
	words := strings.Split(s, " ")
	mapcount := make(map[string]int, 10)
	for _, key := range words {
		v, ok := mapcount[key]
		if ok {
			mapcount[key] = v + 1
		} else {
			mapcount[key] = 1
		}
	}
	fmt.Println(mapcount)

}
