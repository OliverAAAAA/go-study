package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性和JSON序列化

//Student 学生
//参数首字母大写，相当于java的public 小写相当于private
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"sex"`
	Name   string `json:"name"`
}

//Class 班级
type Class struct {
	Title    string    `json:"title"`
	Students []Student `json:"ss"`
}

func newStudent(id int, name, sex string) Student {
	return Student{
		ID:     id,
		Gender: sex,
		Name:   name,
	}
}

func getSex(i int) string {
	if i%2 == 0 {
		return "female"
	} else {
		return "male"
	}
}

func main() {

	c1 := &Class{
		Title:    "101",
		Students: make([]Student, 0, 20),
	}

	for i := 0; i < 10; i++ {
		c1.Students = append(c1.Students, newStudent(i, fmt.Sprintf("stu%2d", i), getSex(i)))
	}

	//结构体 -> json
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("%s\n", data)
	//反序列化 json -> 结构体

	c2 := &Class{}
	json.Unmarshal([]byte(data), c2)
	fmt.Printf("%#v\n", c2)
}
