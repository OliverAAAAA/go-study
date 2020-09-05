package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang~",
		"code":    200,
	})
}

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/hello", sayHello)

	//返回json
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "oliver",
			"age":  25,
		}
		c.JSON(http.StatusOK, data)
	})

	//返回struct
	r.GET("/struct", func(c *gin.Context) {
		type s struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		data := s{
			Name: "Oliver",
			Age:  25,
		}
		c.JSON(http.StatusOK, data) //JSON序列化
	})
	//获取get参数
	r.GET("/query", func(c *gin.Context) {
		//获取请求参数
		age := c.Query("age")
		//取不到就默认值
		name := c.DefaultQuery("name", "Oliver")
		//取不到就返回bool
		sex, ok := c.GetQuery("sex")
		if !ok {
			sex = "1"
		}
		c.JSON(http.StatusOK, gin.H{
			"age":  age,
			"name": name,
			"sex":  sex,
		})

	})
	r.LoadHTMLFiles("33-gin/login.html", "33-gin/index.html")
	//获取form参数
	r.GET("/login", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", nil)
	})

	//获取form参数
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		username2, ok := c.GetPostForm("username")
		if ok {
			fmt.Println(username2)
		}
		password := c.DefaultPostForm("password", "123456")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	//获取restful
	r.GET("/restful/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"age":  age,
			"name": name,
		})
	})
	//结构体接收参数
	r.GET("/structParam", func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"u": user.Username,
				"p": user.Password,
			})

		} else {
			fmt.Printf("convert param error,err:%v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	})

	r.Run(":9090")
}
