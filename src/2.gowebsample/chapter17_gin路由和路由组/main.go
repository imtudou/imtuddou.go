package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	// 1. 访问 /index 的 GET/POST/PUT/DELETA 请求会走下面逻辑
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Path":   "index",
			"Method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Path":   "index",
			"Method": "POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Path":   "index",
			"Method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Path":   "index",
			"Method": "DELETE",
		})
	})

	// 2. any 请求方法的大集合
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(200, gin.H{"Path": "user", "Method": c.Request.Method})
		case "POST":
			c.JSON(200, gin.H{"Path": "user", "Method": c.Request.Method})
		case "PUT":
			c.JSON(200, gin.H{"Path": "user", "Method": c.Request.Method})
		case "DELETE":
			c.JSON(200, gin.H{"Path": "user", "Method": c.Request.Method})
		}
	})

	// 3. no route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "NoRoute", "Method": c.Request.Method})
	})

	// 4. 路由组
	//r.GET("/shop/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	//})
	//r.GET("/shop/item", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/shop/item"})
	//})

	// 路由组的组 多用于区分不同的业务线或API版本
	// 把公用的前缀提取出来，创建一个路由组
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
		})
		shopGroup.GET("/item", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/item"})
		})

		// 嵌套路由组
		product := shopGroup.Group("item")
		product.GET("/food", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/item/food"})
		})
	}

	r.Run(":9090")

	//Gin框架中的路由使用的是httprouter这个库。
	//
	//其基本原理就是构造一个路由地址的前缀树。
}
