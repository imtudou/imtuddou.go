package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取 query 提交的参数
// GET请求 URL ?后面的是querystring参数
// key=value格式，多个key-value用 & 连接
// eq:  /web/query=小王子&age=18
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {

		// 获取querystring参数
		name := c.Query("name")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
