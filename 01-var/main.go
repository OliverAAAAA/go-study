package main

import (
	"fmt"
	"math"
)

const (
	pi = 3.14
	e = 2.7
)

const (
	n1 = iota
	n2
	n3 =100
	n4 =iota
	n5
)
const (
	n6 = iota

)

const (
	_  = iota
	KB = 1 << (10 * iota)// 1<<10  ->   2^10
	MB = 1 << (10 * iota)// 1<<20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	m1, m2 = iota + 1, iota + 2 //1,2
	m3, m4                      //2,3
	m5, m6                      //3,4
)
func main() {

	var (
		a string
		b int
	)
	a = "嗨"
	b = 1
	fmt.Println(a, b)

	m := "abc"
	fmt.Println(m)

	fmt.Println(n1,n2,n3,n4,n5,n6)
	fmt.Println(KB,MB,GB,TB,PB)
	fmt.Println(math.Pow(2,10))
	fmt.Println(m1,m2,m3,m4,m5,m6)

	c := b
	c = 3
	fmt.Println(b,c)
}
