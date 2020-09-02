package main

import (
	"fmt"
)

func main() {
	var name []string
	var age []int
	var flag = []bool{true, false}

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(flag)

	a := [5]int{10, 11, 12, 13, 14}
	//切片
	b := a[1:4]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//切片的切片
	c := b[0:len(b)]
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	//make构建切片
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(len(d))
	fmt.Println(cap(d))

	var e []int     //声明int切片
	var f = []int{} //声明并初始化切片
	var g = make([]int, 0)

	if e == nil {
		fmt.Println("e是一个nil")
	}
	if f != nil {
		fmt.Println("f不是一个nil")
	}
	if g != nil {
		fmt.Println("g不是一个nil")
	}
	fmt.Println(e, len(e), cap(e))
	fmt.Println(f, len(f), cap(f))
	fmt.Println(g, len(g), cap(g))

	//复制的切片，底层公用一个数组
	h := make([]int, 3)
	i := h
	i[0] = 100
	fmt.Println(h)
	fmt.Println(i)

	//遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	fmt.Println("-------")
	for index, value := range a {
		fmt.Println(index, value)
	}

	//append 切片的扩容
	var j []int
	j = append(j, 10)
	fmt.Println(j)
	for i := 0; i < 10; i++ {
		//单条追加
		//j = append(j,i)
		//多条追加
		j = append(j, i, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", j, len(j), cap(j), j)
	}
	//追加切片
	j = append(j, b...)
	fmt.Println(j)

	//切片的copy
	k := []int{1,2,3,4,5}
	l := make([]int,5,5)
	m := k
	copy(l,k)
	k[0] = 100
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

	//切片的删除
	o := []string{"上海","武汉","深圳","北京"}
	o  = append(o[0:2],o[3:]...)
	fmt.Println(o)


}
