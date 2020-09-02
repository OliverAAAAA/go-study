package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello := func() {
		fmt.Println("匿名函数1")
	}
	sayHello()
	//加()立即执行
	func() {
		fmt.Println("匿名函数2")
	}()

	//闭包 = 函数 + 外层变量的引用
	r := a(18)
	r()

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

//定义一个函数，返回值是一个函数
func a(age int) func(){
	name := "Oliver"
	 return func(){
		fmt.Println("hello ",name,",age:",age)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
