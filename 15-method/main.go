package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int
}

//NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func (p Person) setAge(age int) {
	p.age = age
}
func (p *Person) setAge2(age int) {
	p.age = age
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Println(p1.age)
	p1.setAge(18) //不变 25
	fmt.Println(p1.age)
	p1.setAge2(10) //变 10
	fmt.Println(p1.age)
}
