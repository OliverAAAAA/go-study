package main

import "fmt"

//结构体构造函数
type person struct {
	name, city string
	age        int
}

func main() {
	p1 := newPerson("Oliver", "WHC", 25)
	fmt.Printf("%#v", p1)
}

func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		age:  age,
		city: city,
	}
}
