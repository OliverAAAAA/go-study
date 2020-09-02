package main

import "fmt"

// 基于int的自定义类型
type MyInt int

//类型别名
type NewInt  = int

func main() {

	var a MyInt
	fmt.Printf("%T,%v \n" ,a ,a)
	var b NewInt
	fmt.Printf("%T,%v \n" ,b ,b)
}

