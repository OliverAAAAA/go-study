package main

import (
	"fmt"
	g "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
	Hobby string
}

func main() {

	db, err := g.Open("mysql", "root:123456@(localhost)/oliver?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Printf("connect db error,err:%v\n", err)
	}
	//创建表
	db.AutoMigrate(&UserInfo{})

	u1:= UserInfo{
		Id:    0,
		Name:  "Oliver",
		Age:   25,
		Hobby: "coding",
	}
	//插入u1
	//db.Create(&u1)
	fmt.Println(u1)
	var u UserInfo
	//查询表中第一条数据
	db.First(&u)
	fmt.Println(u)
	//更新u
	db.Model(&u).Update("hobby","赚钱")
	//删除u
	db.Delete(&u)
}
