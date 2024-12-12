package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {

		// 跳转到百度
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {

		// 跳转到 /b 对应的酷游处理喊出
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求到了b",
		})
	})

	r.Run(":9090")
}
