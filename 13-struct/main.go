package main

import "fmt"

//定义结构体,结构体是内存对齐的
type person struct {
	name, city string
	age        int8
}

func main() {

	//实例化1
	var p1 person
	p1.age = 25
	p1.name = "oliver"
	p1.city = "Wuhan"
	fmt.Printf("p1= %#v\n", p1)

	//匿名结构体
	var p2 struct {
		name string
		sex  string
	}
	p2.name = "lili"
	p2.sex = "female"
	fmt.Printf("p2= %#v\n", p2)

	//结构体指针
	p3 := new(person)
	fmt.Printf("%T", p3)
	(*p3).name = "tom"
	(*p3).age = 25
	(*p3).city = "csc"
	p3.name = "jet"
	p3.age = 18
	p3.city = "bjc"
	fmt.Printf("p3= %#v\n", p3)

	//取结构体地址
	p4 := &person{}
	p4.name = "catty"
	p4.city = "hk"
	p4.age = 12
	fmt.Printf("p4= %#v\n", p4)

}
