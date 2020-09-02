package main

import (
	"fmt"
	"math"
)



func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF

	var age uint8
	println(age)

	fmt.Printf("%.f \n", math.Pi)  // FF

	var flag  bool
	println(flag)

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	multiStr :=
`
第一行
第二行
第三行
`
	println(multiStr)
}
