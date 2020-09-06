package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"https://baidu.com")
	})

	// /a ->/b
	r.GET("/a", func(c *gin.Context) {
		//把uri请求修改
		c.Request.URL.Path = "/b"
		//继续处理后续
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"bbbbbbb",
		})
	})

	r.Run(":9090")
}