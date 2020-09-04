package main

import (
	"fmt"
	"reflect"
)

//结构体反射

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(s.Name, msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(s.Name, msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := student{
		Name:  "Oliver",
		Score: 120,
	}

	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v,kind:%v\n", t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		filedObj := t.Field(i)
		fmt.Printf("name:%v,type:%v,tag:%v\n", filedObj.Name, filedObj.Type, filedObj.Tag)
		fmt.Println(filedObj.Tag.Get("json"), filedObj.Tag.Get("ini"))
	}
	filedObj, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v,type:%v,tag:%v\n", filedObj.Name, filedObj.Type, filedObj.Tag)
	}

	printMethod(stu1)
}
