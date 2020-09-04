package main

import (
	"fmt"
	"reflect"
)

type cat struct {
}

func reflectType(x interface{}) {
	//反射
	v := reflect.TypeOf(x)
	fmt.Println(v, v.Name(), v.Kind())
	fmt.Printf("%T\n", v)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v, v.Kind())
	t := v.Kind()
	switch t {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("这里:%v,%T\n", ret, ret)
	}
	//获取原类型的变量
}
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a float32 = 3.1415926
	var b int8 = 18
	reflectType(a)
	reflectType(b)
	var c cat
	reflectType(c)
	d := make([]int, 10)
	f := make([]string, 10)
	reflectType(d)
	reflectType(f)
	fmt.Println("~~~~~~~~~~~~~~~")
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(d)
	reflectValue(f)
	fmt.Println("*************")
	var setV int64 = 18
	reflectSetValue1(&setV) //传值会panic，传指针无效
	fmt.Println(setV)
	reflectSetValue2(&setV) //Elem() 操作指针
	fmt.Println(setV)
}
