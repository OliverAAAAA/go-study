package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
GIN中间件
Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，
中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
*/
func filter(c *gin.Context) {
	fmt.Println("time in")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("use time %v\n", cost)
	fmt.Println("time end")
}

func handelRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello 中间件",
	})
}

//闭包写法
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//查询数据
	return func(c *gin.Context) {
		if doCheck {
			//写业务
			c.Next()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(filter)
	r.GET("/index", handelRequest)

	r.Run(":9090")
}
