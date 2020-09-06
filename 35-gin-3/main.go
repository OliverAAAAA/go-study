package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由与路由组

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "get",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "post",
			})

		}
	})
	//404处理
	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"code": 404,
	//	})
	//})

	//路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"uri": "/video/index"})

		})
		videoGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"uri": "/video/login"})
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"uri": "/shop/index"})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"uri": "/shop/cart"})
		})
		shopGroup.GET("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"uri": "/shop/checkout"})
		})

		//路由嵌套
		shopGroupV2 := shopGroup.Group("/v2")
		{
			shopGroupV2.GET("/getPrice", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"uri": "/shop/v2/getPrice"})
			})
		}
	}

	r.Run(":9090")
}
