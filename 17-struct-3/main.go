package main

import "fmt"

//结构体匿名字段-不常用
//Person 结构体Person类型
type Person struct {
	string
	int
}
//嵌套结构体
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

//结构体的继承
type Animal struct {
	name string
}

func (a *Animal) move(){
	fmt.Printf("%s会动~\n",a.name)
}

type dog struct {
	 Feet int
	 *Animal
}

func (d *dog) wang(){
	fmt.Printf("%s会叫~\n",d.name)
}



func main() {
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	fmt.Println(p1.string, p1.int) //北京 18

	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.Address.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	dog := &dog{
		Feet: 4,
		Animal :&Animal{
			name : "dog",
		},
	}
	dog.move()
	dog.wang()
}
