package main

import "fmt"

func main() {
	var a *int //-->  a is nil  ，没有初始化
	a = new(int)
	*a = 100
	fmt.Println(*a)

	b := make(map[string]int, 8)
	b["oliver"] = 100
	fmt.Println(b)
}
